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

const (
	minRaidSize = 10
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
		"raid":   newBot.raidHandler,
		"buy":    newBot.buyHandler,
		// "use":    newBot.useHandler,
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

func (tb *TelegramBot) SendTextMessage(ctx context.Context, chatId int64, text string) *models.Message {
	m, err := tb.bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    chatId,
		Text:      text,
		ParseMode: models.ParseModeMarkdown,
	})
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}

	return m
}

func (tb *TelegramBot) defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	tb.SendTextMessage(
		ctx,
		update.Message.Chat.ID,
		prompt.Default(),
	)
}

func (tb *TelegramBot) helpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	tb.SendTextMessage(
		ctx,
		update.Message.Chat.ID,
		prompt.Help(),
	)
}

func (tb *TelegramBot) jungleHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	jungle := mock.Jungles[0]

	tb.SendTextMessage(
		ctx,
		update.Message.Chat.ID,
		prompt.JungleFromModel(&jungle),
	)
}

func (tb *TelegramBot) raidHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	//TODO:
	// check whether the user currently in raid
	// find monkeys with with similar army size and list them
	// 		if none, list other available jungles
	//		if none, list the Heaven Jungle

	jungle := mock.Jungles[0]

	if jungle.Monkeys < minRaidSize {
		tb.SendTextMessage(
			ctx,
			update.Message.Chat.ID,
			prompt.RaidUnavailable(),
		)
	}

	tb.SendTextMessage(
		ctx,
		update.Message.Chat.ID,
		prompt.RaidList(mock.Jungles),
	)
}

func (tb *TelegramBot) buyHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	tb.SendTextMessage(
		ctx,
		update.Message.Chat.ID,
		prompt.Buy(),
	)
}

// TODO: use functionality
func (tb *TelegramBot) useHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	jungle := mock.Jungles[0]

	tb.SendTextMessage(
		ctx,
		update.Message.Chat.ID,
		prompt.JungleFromModel(&jungle),
	)
}
