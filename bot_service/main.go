package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/babyress/telegram-bot/config"
	"github.com/babyress/telegram-bot/handlers"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var adminUserIDs []int64
var appConfig *config.Config
var telegramBot *bot.Bot

func main() {
	appConfig = config.LoadConfig()

	// Start the health check server
	go startHealthCheckServer()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Println("Received termination signal. Shutting down...")
		cancel()
	}()

	userHandler := handlers.NewUserHandler(appConfig, appConfig.TelegramGroupURL)
	opts := []bot.Option{
		bot.WithDefaultHandler(func(ctx context.Context, b *bot.Bot, update *models.Update) {
			handleMessage(ctx, b, update, userHandler)
		}),
	}
	botToken := appConfig.BotToken
	b, err := bot.New(botToken, opts...)
	if err != nil {
		log.Fatal(err)
	}
	// Store the bot instance in the global variable
	telegramBot = b
	fmt.Println("Bot Service started")
	b.Start(ctx)
}

func handleMessage(ctx context.Context, b *bot.Bot, update *models.Update, userHandler *handlers.UserHandler) {
	if update.Message == nil {
		return
	}
	user := handlers.User{
		ID:        update.Message.From.ID,
		FirstName: update.Message.From.FirstName,
		LastName:  update.Message.From.LastName,
		Username:  update.Message.From.Username,
		IsAdmin:   handlers.IsAdmin(update.Message.From.ID, adminUserIDs),
	}

	userHandler.HandleMessage(ctx, b, update, user)
}

// SendCodeRequest defines the structure for the sendCode endpoint request
type SendCodeRequest struct {
	ChatID  int64  `json:"chatId"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

// handleSendCode handles POST requests to /sendCode
func handleSendCode(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse request body
	var req SendCodeRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Error parsing request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate request
	if req.ChatID == 0 {
		http.Error(w, "ChatID is required", http.StatusBadRequest)
		return
	}

	// Prepare message text
	var messageText string
	if req.Message != "" {
		// Use provided message if available
		messageText = req.Message
	} else if req.Code != "" {
		// Use voucher code template if code is provided
		messageText = fmt.Sprintf("Hi! You have new voucher code %s. Please claim here http://gooogle.com", req.Code)
	} else {
		// Fall back to default template if neither message nor code is provided
		messageText = "Hi! You have new voucher code. Please claim here http://gooogle.com"
	}

	// Send message to Telegram user
	ctx := context.Background()
	_, err = telegramBot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: req.ChatID,
		Text:   messageText,
	})

	if err != nil {
		log.Printf("Error sending message to Telegram: %v", err)
		http.Error(w, "Error sending message to Telegram: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Code sent successfully",
	})
}

// startHealthCheckServer starts a simple HTTP server for health checks
func startHealthCheckServer() {
	// Register health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy","service":"bot_service"}`))
	})

	// Register sendCode endpoint
	http.HandleFunc("/sendCode", handleSendCode)

	port := ":8080"
	log.Printf("HTTP server started on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
