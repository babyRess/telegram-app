package handlers

import (
	"fmt"

	"github.com/go-telegram/bot/models"
)

type Environment string

type ButtonType string

const (
	ButtonTypeRocket    ButtonType = "rocket"
	ButtonTypeBank      ButtonType = "bank"
	ButtonTypeMoneyTree ButtonType = "money_tree"
)

type ButtonOptions struct {
	Environment Environment
}

func CreateGameButtons(groupURL string, isDevMode bool) *models.InlineKeyboardMarkup {
	joinGroupButton := models.InlineKeyboardButton{
		Text: "Join Telegram Group",
		URL:  groupURL,
	}
	buttons := [][]models.InlineKeyboardButton{
		{joinGroupButton},
		{CreateGameButton("Rocket", ButtonTypeRocket, isDevMode)},
		{CreateGameButton("Bank", ButtonTypeBank, isDevMode)},
		{CreateGameButton("Money Tree", ButtonTypeMoneyTree, isDevMode)},
	}
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: buttons,
	}
}

// createGameButton creates a single game button
func CreateGameButton(name string, gameType ButtonType, isDevMode bool) models.InlineKeyboardButton {
	envParam := ""
	if isDevMode {
		envParam = "&gameEnv=local"
	}

	var devMode string
	if isDevMode {
		devMode = "(DEV)"
	} else {
		devMode = ""
	}
	url := "https://api.g1388.makethatold.com/promobot/lobby/?gameIdentifier=" + string(gameType) + envParam

	return models.InlineKeyboardButton{
		Text: fmt.Sprintf("(%s) Try Your Luck! %s", name, devMode),
		WebApp: &models.WebAppInfo{
			URL: url,
		},
	}
}

func CreateGameButtonWithCustomText(name string, gameType ButtonType, isDevMode bool, customText string) models.InlineKeyboardButton {
	envParam := ""
	if isDevMode {
		envParam = "&gameEnv=local"
	}

	var devMode string
	if isDevMode {
		devMode = "(DEV)"
	} else {
		devMode = ""
	}
	fmt.Println("https://api.g1388.makethatold.com/promobot/lobby/?gameIdentifier=" + string(gameType) + "&rewardType=" + customText + envParam)

	return models.InlineKeyboardButton{
		Text: fmt.Sprintf("%s with %s %s", name, customText, devMode),
		WebApp: &models.WebAppInfo{
			URL: "https://api.g1388.makethatold.com/promobot/lobby/?gameIdentifier=" + string(gameType) + "&rewardType=" + customText + envParam,
		},
	}
}
