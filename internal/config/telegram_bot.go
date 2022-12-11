package config

type TelegramBotConfig interface {
	TelegramToken() string
	Offset() int
	Timeout() int
}

type config struct {
	telegramToken string
	offset        int
	timeout       int
}

func NewConfig() (TelegramBotConfig, error) {
	return &config{
		telegramToken: "token",
	}, nil
}

func (c *config) TelegramToken() string {
	return c.telegramToken
}

func (c *config) Offset() int {
	return c.offset
}

func (c *config) Timeout() int {
	return c.timeout
}
