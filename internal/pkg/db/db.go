package db

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Query struct {
	Name     string
	QueryRaw string
}

type DB struct {
	pool *pgxpool.Pool
}

func (db *DB) GetContext(ctx context.Context, dest any, q Query, args ...any) error {
	return pgxscan.Get(ctx, db.pool, dest, q.QueryRaw, args...)
}

func (db *DB) SelectContext(ctx context.Context, dest any, q Query, args ...any) error {
	return pgxscan.Select(ctx, db.pool, dest, q.QueryRaw, args...)
}

func (db *DB) ExecContext(ctx context.Context, q Query, args ...any) (pgconn.CommandTag, error) {
	return db.pool.Exec(ctx, q.QueryRaw, args...)
}

func (db *DB) QueryContext(ctx context.Context, q Query, args ...any) (pgx.Rows, error) {
	return db.pool.Query(ctx, q.QueryRaw, args...)
}

func (db *DB) QueryRowContext(ctx context.Context, q Query, args ...any) pgx.Row {
	return db.pool.QueryRow(ctx, q.QueryRaw, args...)
}

func (db *DB) Close() {
	db.pool.Close()
}
