package repository

import "github.com/lardira/monking/internal/domain"

type UserRepository interface {
	GetAll() ([]domain.User, error)
	GetByID(string) (*domain.User, error)
}
