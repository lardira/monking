package contextkeys

import (
	"context"

	"github.com/lardira/monking/internal/domain"
)

type ContextKey string

const (
	ContextKeyUser ContextKey = "user"
)

func UserFromContext(ctx context.Context) (*domain.User, bool) {
	val := ctx.Value(ContextKeyUser)
	user, ok := val.(*domain.User)
	return user, ok
}
