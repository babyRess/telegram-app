package models
import (
	"time"
	"github.com/go-telegram/bot/models"
)
type User struct {
	ID        int64  `bson:"id"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name,omitempty"`
	Username  string `bson:"username,omitempty"`
	IsAdmin   bool   `bson:"is_admin"`
}
type LogData struct {
	UserInfo *models.User     `bson:"user_info"`
	Date     time.Time        `bson:"date"`
	Location *models.Location `bson:"location,omitempty"`
}
type ButtonType string
const (
	ButtonTypeRocket    ButtonType = "rocket"
	ButtonTypeBank      ButtonType = "bank"
	ButtonTypeMoneyTree ButtonType = "money_tree"
)
type Environment string
const (
	EnvironmentProduction  Environment = "production"
	EnvironmentDevelopment Environment = "local"
)
