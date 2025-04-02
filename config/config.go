package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken         string
	ChannelID        int64
	DefaultAppText   string
	DefaultAppURL    string
	MongoURI         string
	Database         string
	TelegramGroupURL string
	MiniAppURLs      map[string]string

	// Voucher configuration
	VoucherButtonText string
	VoucherTypeParam  string

	// KYC configuration
	KYCButtonText string
	KYCTypeParam  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables.")
	}
	botToken := getEnv("TELEGRAM_BOT_TOKEN", "")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set")
	}
	channelIDStr := getEnv("TELEGRAM_CHANNEL_ID", "-1002254268367")
	channelID, err := strconv.ParseInt(channelIDStr, 10, 64)
	if err != nil {
		log.Printf("Warning: Invalid TELEGRAM_CHANNEL_ID, using default: %v", err)
		channelID = -1002254268367
	}
	// mongoURI := getEnv("MONGO_URI", "mongodb+srv://datnt13022:asdasd11@cluster0.juehd.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	// database := getEnv("MONGO_DATABASE", "telegram_bot")
	telegramGroupURL := getEnv("TELEGRAM_GROUP_URL", "https://t.me/mkt_royalgame")
	miniAppURLs := make(map[string]string)
	miniAppURLs["rocket"] = getEnv("MINIAPP_URL_ROCKET", "https://api.g1388.makethatold.com/promobot/lobby/?gameIdentifier=rocket")
	miniAppURLs["bank"] = getEnv("MINIAPP_URL_BANK", "https://api.g1388.makethatold.com/promobot/lobby/?gameIdentifier=bank")
	miniAppURLs["money_tree"] = getEnv("MINIAPP_URL_MONEY_TREE", "https://api.g1388.makethatold.com/promobot/lobby/?gameIdentifier=money_tree")

	// Load voucher and KYC configurations
	voucherButtonText := getEnv("VOUCHER_BUTTON_TEXT", "Redeem Voucher")
	voucherTypeParam := getEnv("VOUCHER_TYPE_PARAM", "redeemcode")

	kycButtonText := getEnv("KYC_BUTTON_TEXT", "Complete KYC")
	kycTypeParam := getEnv("KYC_TYPE_PARAM", "kyc")

	return &Config{
		BotToken:         botToken,
		ChannelID:        channelID,
		DefaultAppText:   getEnv("DEFAULT_APP_TEXT", "Try Your Luck!🍀"),
		DefaultAppURL:    getEnv("DEFAULT_APP_URL", "https://api.g1388.makethatold.com/promobot/lobby/?gameIdentifier=bank"),
		TelegramGroupURL: telegramGroupURL,
		MiniAppURLs:      miniAppURLs,

		VoucherButtonText: voucherButtonText,
		VoucherTypeParam:  voucherTypeParam,
		KYCButtonText:     kycButtonText,
		KYCTypeParam:      kycTypeParam,
	}
}
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
func GetAdminUserIDs() []int64 {
	adminIDsStr := getEnv("ADMIN_USER_IDS", "")
	if adminIDsStr == "" {
		return []int64{}
	}
	adminIDsStrSlice := strings.Split(adminIDsStr, ",")
	adminIDs := make([]int64, 0, len(adminIDsStrSlice))
	for _, idStr := range adminIDsStrSlice {
		id, err := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64)
		if err != nil {
			log.Printf("Warning: Invalid admin user ID %s: %v", idStr, err)
			continue
		}
		adminIDs = append(adminIDs, id)
	}
	return adminIDs
}
