package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type BotSettings struct {
	Ctx    context.Context
	Bot    *bot.Bot
	Update *models.Update
}

func NewBotSettings(ctx context.Context, b *bot.Bot, update *models.Update) BotSettings {
	return BotSettings{
		Ctx:    ctx,
		Bot:    b,
		Update: update,
	}
}
