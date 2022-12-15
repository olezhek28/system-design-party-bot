package telegram

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/mock_telegram_client.go -package=mocks . Client

import (
	"log"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/config"
)

type Client interface {
	Start() tgBotAPI.UpdatesChannel
	Send(msg tgBotAPI.MessageConfig) error
	Request(c tgBotAPI.Chattable) (*tgBotAPI.APIResponse, error)
}

type client struct {
	tgBot *tgBotAPI.BotAPI
	cfg   config.TelegramBotConfig
}

func NewClient(tgBot *tgBotAPI.BotAPI, cfg config.TelegramBotConfig) Client {
	return &client{
		tgBot: tgBot,
		cfg:   cfg,
	}
}

func (c *client) Start() tgBotAPI.UpdatesChannel {
	log.Printf("Authorized on account %s", c.tgBot.Self.UserName)

	u := tgBotAPI.NewUpdate(c.cfg.Offset())
	u.Timeout = c.cfg.Timeout()

	return c.tgBot.GetUpdatesChan(u)
}

func (c *client) Send(msg tgBotAPI.MessageConfig) error {
	_, err := c.tgBot.Send(msg)
	return err
}

func (c *client) Request(callback tgBotAPI.Chattable) (*tgBotAPI.APIResponse, error) {
	_, err := c.tgBot.Request(callback)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
