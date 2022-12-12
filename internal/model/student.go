package model

import "time"

type Student struct {
	ID               int64     `db:"id"`
	FirstName        string    `db:"first_name"`
	LastName         string    `db:"last_name"`
	TelegramChatID   int64     `db:"telegram_chat_id"`
	TelegramUsername string    `db:"telegram_username"`
	CreatedAt        time.Time `db:"created_at"`
}
