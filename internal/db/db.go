package db

import (
	"database/sql"
	"errors"
)

var (
	ErrUserNotFound error = errors.New("user not found")
)

func NullStringToPtr(s *sql.NullString) *string {
	if s.Valid {
		return &s.String
	}

	return nil
}
