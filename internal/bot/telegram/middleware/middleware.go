package middleware

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/lardira/monking/internal/contextkeys"
	"github.com/lardira/monking/internal/domain"
)

func UserAuth(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, bot *bot.Bot, update *models.Update) {
		// TODO: auth
		// check if user has account on the channel (e.g. if it has userId of telegram in the db)
		//	no: transfer to registration/loing, should be one command
		//	yes: next handler can be called, context is filled with user entity
		ctx = context.WithValue(ctx, contextkeys.ContextKeyUser, domain.User{})
		next(ctx, bot, update)
	}
}
