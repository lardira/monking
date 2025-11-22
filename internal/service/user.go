package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/lardira/monking/internal/db"
	"github.com/lardira/monking/internal/db/repository"
	"github.com/lardira/monking/internal/domain"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) Create(telegramId *string, discordId *string) (*domain.User, error) {
	id := uuid.NewString()
	u, err := s.userRepository.Create(id, telegramId, discordId)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *UserService) FindOrCreateByTelegramID(telegramId string) (*domain.User, error) {
	user, err := s.userRepository.GetByTelegramID(telegramId)
	if err != nil {
		if errors.Is(err, db.ErrUserNotFound) {
			return s.Create(&telegramId, nil)
		}

		return nil, err
	}

	return user, nil
}
