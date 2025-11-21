package sqlite

import "github.com/lardira/monking/internal/domain"

type UserRepository struct {
}

func (ur *UserRepository) GetAll() ([]domain.User, error) {
	return nil, nil
}

func (ur *UserRepository) GetByID(id string) (*domain.User, error) {
	return nil, nil
}
