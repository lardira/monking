package telegram

import (
	"context"
	"errors"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/lardira/monking/internal/bot/telegram/prompt"
	"github.com/lardira/monking/internal/mock"
)

type TelegramBot struct {
	bot        *bot.Bot
	handlerIDs []string
}

func New(apiToken string) (*TelegramBot, error) {
	if apiToken == "" {
		return nil, errors.New("invalid telegram api token: cannot be empty")
	}

	newBot := &TelegramBot{}

	b, err := bot.New(
		apiToken,
		bot.WithDefaultHandler(newBot.defaultHandler),
	)
	if err != nil {
		return nil, err
	}
	newBot.bot = b

	handlers := map[string]func(context.Context, *bot.Bot, *models.Update){
		"start":  newBot.jungleHandler,
		"jungle": newBot.jungleHandler,
		"help":   newBot.helpHandler,
	}
	for command, handler := range handlers {
		newBot.bot.RegisterHandler(
			bot.HandlerTypeMessageText,
			command,
			bot.MatchTypeCommand,
			handler,
		)
	}

	return newBot, nil
}

func (tb *TelegramBot) Start(ctx context.Context) {
	log.Println("Starting telegram bot...")

	tb.bot.Start(ctx)
}

func (tb *TelegramBot) SendTextMessage(ctx context.Context, chatId int64, text string) (*models.Message, error) {
	m, err := tb.bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    chatId,
		Text:      text,
		ParseMode: models.ParseModeMarkdown,
	})
	return m, err
}

func (tb *TelegramBot) defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := tb.SendTextMessage(
		ctx,
		update.Message.Chat.ID,
		prompt.Default(),
	)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func (tb *TelegramBot) helpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := tb.SendTextMessage(
		ctx,
		update.Message.Chat.ID,
		prompt.Help(),
	)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func (tb *TelegramBot) jungleHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	jungle := mock.Jungles[0]

	_, err := tb.SendTextMessage(
		ctx,
		update.Message.Chat.ID,
		prompt.JungleFromModel(&jungle),
	)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
