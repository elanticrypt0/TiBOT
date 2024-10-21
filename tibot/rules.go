package tibot

import "github.com/go-telegram/bot"

type Rule struct {
	HandlerType bot.HandlerType
	Regex       string
	HandlerFunc bot.HandlerFunc
}
