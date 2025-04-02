package handlers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// User represents a Telegram user with additional information
type User struct {
	ID        int64
	FirstName string
	LastName  string
	Username  string
	IsAdmin   bool
}

// DatabaseInterface defines the interface for database operations
type DatabaseInterface interface {
	GetAdminUserIDs(ctx context.Context) ([]int64, error)
	GetAllUserIDs(ctx context.Context) ([]int64, error)
	LogUserActivity(ctx context.Context, logData interface{}) error
	CreateUser(ctx context.Context, user *models.User, isAdmin bool) error
}

// ConfigInterface defines the interface for configuration
type ConfigInterface interface {
	GetChannelID() int64
	GetTelegramGroupURL() string
	GetMiniAppURL(gameType string) string
	GetDefaultAppText() string
	GetDefaultAppURL() string
}

// SendMessageToChannel sends a message to the configured channel
func SendMessageToChannel(ctx context.Context, b *bot.Bot, channelID int64) {
	fmt.Println("Sending message to channel")

	// Create inline button
	inlineKeyboard := [][]models.InlineKeyboardButton{
		{{Text: "Try Your Luck!🍀", URL: "https://t.me/tetleleksmgv_bot?startapp=gameapp&gameIdentifier=bank"}},
	}

	_, err := b.SendPhoto(ctx, &bot.SendPhotoParams{
		ChatID: channelID,
		Caption: `🚀 NO WINNING? NO PROBLEM! 🚀
💰💰FREE MONEY JUST FOR PLAYING!💰💰

🐹 Dear Royaler,

👉 We have added many NEW games to our gamification services check the Game list [here](https://t.me/tetleleksmgv_bot)

😉 If you decide to keep your tokens instead of CLAIMING them, you'll be (very) pleasantly surprised!

👍 Our game will boost the prizes and continue to give you FREE REWARDS!

🧡 Stay tuned! We've prepared more exciting features!!!`,
		ReplyMarkup: &models.InlineKeyboardMarkup{InlineKeyboard: inlineKeyboard},
		ParseMode:   "Markdown", // or "HTML"
		Photo: &models.InputFileString{
			Data: "https://staging-acegames.s3-ap-southeast-1.amazonaws.com/uploads/telebots.jpeg",
		},
	})
	if err != nil {
		log.Printf("Failed to send photo: %v", err)
	}
}

// IsAdmin checks if a user ID is in the admin list
func IsAdmin(userID int64, adminIDs []int64) bool {
	for _, id := range adminIDs {
		if id == userID {
			return true
		}
	}
	return false
}

// CreateLogData creates a LogData struct from a Telegram user
func CreateLogData(user *models.User, location *models.Location) interface{} {
	return struct {
		UserInfo *models.User     `bson:"user_info"`
		Date     time.Time        `bson:"date"`
		Location *models.Location `bson:"location,omitempty"`
	}{
		UserInfo: user,
		Date:     time.Now(),
		Location: location,
	}
}
