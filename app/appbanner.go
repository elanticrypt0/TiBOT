package app

import (
	"fmt"

	"github.com/go-telegram/bot/models"
)

func AppBanner(botinfo *models.User) {

	botInfo := fmt.Sprintf("BotInfo : %#v", botinfo)

	fmt.Println("")
	fmt.Println("::  TiBOT ::")
	fmt.Println("::   ON   ::")
	fmt.Println("::        ::")
	fmt.Println(botInfo)
	fmt.Println("")
}
