package sqlite

import "github.com/lardira/monking/internal/domain"

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) GetAll() ([]domain.User, error) {
	return nil, nil
}

func (ur *UserRepository) GetByID(id string) (*domain.User, error) {
	return nil, nil
}

func (ur *UserRepository) GetByTelegramID(string) (*domain.User, error) {
	return nil, nil
}

func (ur *UserRepository) Create(id string, telegramId *string, discordId *string) (*domain.User, error) {
	return nil, nil
}

func (ur *UserRepository) Update(id string, telegramId *string, discordId *string) (*domain.User, error) {
	return nil, nil
}
