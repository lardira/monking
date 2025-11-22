package sqlite

import (
	"database/sql"
	"errors"

	"github.com/lardira/monking/internal/db"
	"github.com/lardira/monking/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetByID(id string) (*domain.User, error) {
	var u domain.User
	var sqlTgID, sqlDisID sql.NullString

	query := `SELECT id, telegram_id, discord_id FROM users WHERE id=$1`
	err := ur.db.QueryRow(query, id).Scan(
		&u.ID,
		&sqlTgID,
		&sqlDisID,
	)
	if err != nil {
		return nil, err
	}

	u.TelegramID = db.NullStringToPtr(&sqlTgID)
	u.DiscordID = db.NullStringToPtr(&sqlDisID)
	return &u, nil
}

func (ur *UserRepository) GetByTelegramID(telegramId string) (*domain.User, error) {
	var u domain.User
	var sqlTgID, sqlDisID sql.NullString

	query := `SELECT id, telegram_id, discord_id 
			  FROM users WHERE telegram_id=$1`
	err := ur.db.QueryRow(query, telegramId).Scan(
		&u.ID,
		&sqlTgID,
		&sqlDisID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, db.ErrUserNotFound
		}
		return nil, err
	}

	u.TelegramID = db.NullStringToPtr(&sqlTgID)
	u.DiscordID = db.NullStringToPtr(&sqlDisID)
	return &u, nil
}

func (ur *UserRepository) Create(id string, telegramId *string, discordId *string) (*domain.User, error) {
	query := `INSERT INTO users (id, telegram_id, discord_id)
			  VALUES ($1, $2, $3)`
	if _, err := ur.db.Exec(query, id, telegramId, discordId); err != nil {
		return nil, err
	}

	u := domain.User{
		ID:         id,
		TelegramID: telegramId,
		DiscordID:  discordId,
	}
	return &u, nil
}

func (ur *UserRepository) Update(id string, telegramId *string, discordId *string) (*domain.User, error) {

	u, err := ur.GetByID(id)
	if err != nil {
		return nil, err
	}

	query := `UPDATE users SET telegram_id=$1, discord_id=$2 WHERE id=$3`
	if _, err := ur.db.Exec(query, telegramId, discordId, id); err != nil {
		return nil, err
	}

	u.TelegramID = telegramId
	u.DiscordID = discordId
	return u, nil
}
