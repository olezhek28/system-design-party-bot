package model

import (
	"database/sql"
	"time"
)

type Student struct {
	ID               int64         `db:"id"`
	FirstName        string        `db:"first_name"`
	LastName         string        `db:"last_name"`
	TelegramID       int64         `db:"telegram_id"`
	TelegramUsername string        `db:"telegram_username"`
	Timezone         sql.NullInt64 `db:"timezone"`
	CreatedAt        time.Time     `db:"created_at"`
}

type UpdateStudent struct {
	FirstName        sql.NullString
	LastName         sql.NullString
	TelegramUsername sql.NullString
	Timezone         sql.NullInt64
}
