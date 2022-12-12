package model

import (
	"database/sql"
	"time"
)

type Topic struct {
	ID        int64        `db:"id"`
	Name      string       `db:"name"`
	Link      string       `db:"link"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
