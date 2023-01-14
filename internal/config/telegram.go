package config

import (
	"errors"
	"os"
	"strconv"
)

const (
	telegramTokenEnvName   = "TELEGRAM_TOKEN"
	telegramOffsetEnvName  = "TELEGRAM_OFFSET"
	telegramTimeoutEnvName = "TELEGRAM_TIMEOUT"

	defaultTelegramOffset  = 0
	defaultTelegramTimeout = 60
)

// TelegramBotConfig ...
type TelegramBotConfig interface {
	TelegramToken() string
	Offset() int
	Timeout() int
}

type telegramConfig struct {
	telegramToken string
	offset        int
	timeout       int
}

// GetTelegramBotConfig ...
func GetTelegramBotConfig() (TelegramBotConfig, error) {
	token := os.Getenv(telegramTokenEnvName)
	if len(token) == 0 {
		return nil, errors.New("telegram token not found")
	}

	offsetStr := os.Getenv(telegramOffsetEnvName)
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil || offset == 0 {
		offset = defaultTelegramOffset
	}

	timeoutStr := os.Getenv(telegramTimeoutEnvName)
	timeout, err := strconv.ParseInt(timeoutStr, 10, 64)
	if err != nil || timeout == 0 {
		timeout = defaultTelegramTimeout
	}

	return &telegramConfig{
		telegramToken: token,
		offset:        int(offset),
		timeout:       int(timeout),
	}, nil
}

// TelegramToken ...
func (c *telegramConfig) TelegramToken() string {
	return c.telegramToken
}

// Offset ...
func (c *telegramConfig) Offset() int {
	return c.offset
}

// Timeout ...
func (c *telegramConfig) Timeout() int {
	return c.timeout
}
