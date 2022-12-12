package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Client interface {
	Close() error
	DB() *DB
}

type client struct {
	db *DB
}

func NewClient(ctx context.Context, dsn string) (Client, error) {
	dbc, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &client{
		db: &DB{pool: dbc},
	}, nil
}

func (c *client) Close() error {
	if c != nil {
		if c.db != nil {
			c.db.Close()
		}
	}

	return nil
}

func (c *client) DB() *DB {
	return c.db
}
