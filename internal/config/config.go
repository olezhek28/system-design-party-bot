package config

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/olezhek28/system-design-party-bot/internal/logger"
)

const stgPostfix = "_STG"

func Init(_ context.Context) error {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Info("No .env file found")
	}

	return nil
}

func get(key string, isStgEnv bool) string {
	if isStgEnv {
		return os.Getenv(key + stgPostfix)
	}

	return os.Getenv(key)
}
