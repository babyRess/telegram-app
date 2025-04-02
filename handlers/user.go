package handlers

import (
	"context"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// UserHandler handles user commands
type UserHandler struct {
	config   ConfigInterface
	groupURL string
}

// NewUserHandler creates a new user handler
func NewUserHandler(config ConfigInterface, groupURL string) *UserHandler {
	return &UserHandler{
		config:   config,
		groupURL: groupURL,
	}
}

// HandleMessage handles user messages
func (h *UserHandler) HandleMessage(ctx context.Context, b *bot.Bot, update *models.Update, user User) {
	if update.Message.Text == "/start" {
		h.handleStart(ctx, b, update, user)
	} else if update.Message.Text == "/help" {
		h.handleHelp(ctx, b, update, user)
	} else if update.Message.Text == "/dev" {
		h.handleDev(ctx, b, update, user)
	} else if update.Message.Text == "/start_voucher" {
		h.handleStartVoucher(ctx, b, update, user)
	} else if update.Message.Text == "/start_kyc" {
		h.handleStartKYC(ctx, b, update, user)
	} else if strings.HasPrefix(update.Message.Text, "/start_voucher") {
		h.handleStartVoucher(ctx, b, update, user)
	} else if strings.HasPrefix(update.Message.Text, "/start_kyc") {
		h.handleStartKYC(ctx, b, update, user)
	}
}

// handleStart handles the /start command
func (h *UserHandler) handleStart(ctx context.Context, b *bot.Bot, update *models.Update, user User) {

	// Create game buttons
	inlineKeyboard := CreateGameButtons(h.groupURL, false)

	// Send welcome message with buttons
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Welcome to the game! Please join the telegram group for news and click the miniapp to start playing.",
		ReplyMarkup: inlineKeyboard,
	})

	// Set the chat menu button
	b.SetChatMenuButton(
		ctx, &bot.SetChatMenuButtonParams{
			ChatID: update.Message.Chat.ID,
			MenuButton: &models.MenuButtonWebApp{
				Type: "web_app",
				Text: h.config.GetDefaultAppText(),
				WebApp: models.WebAppInfo{
					URL: h.config.GetDefaultAppURL(),
				},
			}})
}

// handleHelp handles the /help command
func (h *UserHandler) handleHelp(ctx context.Context, b *bot.Bot, update *models.Update, user User) {
	helpMessage := "Available Commands:\n" +
		"/start - Start the bot and get game links\n" +
		"/help - Show this help message\n" +
		"/dev - Show development game links (for testing)\n" +
		"/start_voucher - Go to code redemption flow\n" +
		"/start_kyc - Go to KYC verification flow"

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   helpMessage,
	})
}

// handleDev handles the /dev command
func (h *UserHandler) handleDev(ctx context.Context, b *bot.Bot, update *models.Update, user User) {
	// Create game buttons for development environment

	inlineKeyboard := CreateGameButtons(h.groupURL, true)

	// Send message with development buttons
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Welcome to the game! These are the development links for testing purposes.",
		ReplyMarkup: inlineKeyboard,
	})
}

// handleStartVoucher handles the /start_voucher command
func (h *UserHandler) handleStartVoucher(ctx context.Context, b *bot.Bot, update *models.Update, user User) {
	// Log user activity
	voucherButton := CreateGameButtonWithCustomText("Bank", ButtonTypeBank, false, "redeemcode")
	voucherDevButton := CreateGameButtonWithCustomText("Bank", ButtonTypeBank, true, "redeemcode")

	inlineKeyboard := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{voucherButton, voucherDevButton},
		},
	}

	// Send message with voucher button
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Click to enter bank reward is redeem code:.",
		ReplyMarkup: inlineKeyboard,
	})
}

// handleStartKYC handles the /start_kyc command
func (h *UserHandler) handleStartKYC(ctx context.Context, b *bot.Bot, update *models.Update, user User) {

	// Get KYC URL with type parameter from config
	kycButton := CreateGameButtonWithCustomText("Bank", ButtonTypeBank, false, "kyc")
	kycDevButton := CreateGameButtonWithCustomText("Bank", ButtonTypeBank, true, "kyc")
	// Create KYC button with text from config

	inlineKeyboard := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{kycButton, kycDevButton},
		},
	}

	// Send message with KYC button
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Click to enter bank reward should kyc:.",
		ReplyMarkup: inlineKeyboard,
	})
}
