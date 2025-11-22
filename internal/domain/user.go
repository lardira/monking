package domain

import "github.com/google/uuid"

type User struct {
	ID         string
	TelegramID *string
	DiscordID  *string
}

func (u *User) Valid() bool {
	err := uuid.Validate(u.ID)
	return err == nil
}
