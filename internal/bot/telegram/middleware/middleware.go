package middleware

import (
	"context"
	"log"
	"strconv"

	tgbot "github.com/go-telegram/bot"

	"github.com/go-telegram/bot/models"
	"github.com/lardira/monking/internal/contextkeys"
	"github.com/lardira/monking/internal/service"
)

func NewUserAuth(userService *service.UserService) tgbot.Middleware {
	return func(next tgbot.HandlerFunc) tgbot.HandlerFunc {
		return func(ctx context.Context, bot *tgbot.Bot, update *models.Update) {
			userTelegramId := strconv.FormatInt(update.Message.From.ID, 10)
			user, err := userService.FindOrCreateByTelegramID(userTelegramId)
			if err != nil {
				_, err := bot.SendMessage(ctx, &tgbot.SendMessageParams{
					ChatID:    update.Message.Chat.ID,
					Text:      "Right now i cannot create an account for you \\:\\(",
					ParseMode: models.ParseModeMarkdown,
				})
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}
				return
			}

			ctx = context.WithValue(ctx, contextkeys.ContextKeyUser, user)
			next(ctx, bot, update)
		}
	}
}
