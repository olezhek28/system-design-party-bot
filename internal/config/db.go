package config

import (
	"errors"
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
func GetDBConfig(isStgEnv bool) (DBConfig, error) {
	dsn := get(dsnEnvName, isStgEnv)
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
