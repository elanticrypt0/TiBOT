package app

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	TelegramApiKey string
	AdminUserIdStr string
	AdminUserId    int64
}

func NewConfig() *Config {
	config := &Config{}
	config.LoadConfig()
	return config

}

func (me *Config) LoadConfig() {
	me.TelegramApiKey = os.Getenv("TELEGRAM_API_KEY")
	me.AdminUserIdStr = os.Getenv("ADMIN_USER_ID")
	auID, err := strconv.Atoi(me.AdminUserIdStr)
	if err != nil {
		log.Panic(err)
	}
	me.AdminUserId = int64(auID)
}
