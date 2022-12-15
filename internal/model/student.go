package model

import "time"

type Student struct {
	ID               int64     `db:"id"`
	FirstName        string    `db:"first_name"`
	LastName         string    `db:"last_name"`
	TelegramID       int64     `db:"telegram_id"`
	TelegramUsername string    `db:"telegram_username"`
	CreatedAt        time.Time `db:"created_at"`
}
