package service

import (
	"database/sql"
	"errors"

	"github.com/lardira/monking/internal/db/repository"
	"github.com/lardira/monking/internal/domain"
)

var (
	ErrUserNotFound error = errors.New("user not found")
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
	return nil, nil
}

func (s *UserService) FindById(id string) {

}

func (s *UserService) FindByTelegramID(telegramId string) (*domain.User, error) {
	user, err := s.userRepository.GetByTelegramID(telegramId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return user, nil
}

func (s *UserService) FindOrCreateByTelegramID(telegramId string) (*domain.User, error) {
	user, err := s.userRepository.GetByTelegramID(telegramId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return s.Create(&telegramId, nil)
		}

		return nil, err
	}

	return user, nil
}
