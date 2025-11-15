package telegram

import (
	"context"
	"errors"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type TelegramBot struct {
	bot *bot.Bot
}

func New(apiToken string) (*TelegramBot, error) {
	if apiToken == "" {
		return nil, errors.New("invalid telegram api token: cannot be empty")
	}

	newBot := &TelegramBot{}

	b, err := bot.New(
		apiToken,
		bot.WithDefaultHandler(newBot.handler),
	)
	if err != nil {
		return nil, err
	}

	newBot.bot = b
	return newBot, nil
}

func (tb *TelegramBot) Start(ctx context.Context) {
	log.Println("Starting telegram bot...")

	tb.bot.Start(ctx)
}

func (tb *TelegramBot) handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	},
	)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
