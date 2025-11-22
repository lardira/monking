package repository

import "github.com/lardira/monking/internal/domain"

type UserRepository interface {
	GetByID(string) (*domain.User, error)
	GetByTelegramID(string) (*domain.User, error)
	Create(id string, telegramId *string, discordId *string) (*domain.User, error)
	Update(id string, telegramId *string, discordId *string) (*domain.User, error)
}
