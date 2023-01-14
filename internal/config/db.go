package config

import (
	"errors"
	"os"
)

const (
	dsnEnvName = "DB_DSN"
)

// DBConfig ...
type DBConfig interface {
	DSN() string
}

type dbConfig struct {
	dsn string
}

// GetDBConfig ...
func GetDBConfig() (DBConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("db dsb not found")
	}

	return &dbConfig{
		dsn: dsn,
	}, nil
}

// DSN ...
func (cfg *dbConfig) DSN() string {
	return cfg.dsn
}
