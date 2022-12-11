package telegram

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/mock_telegram_client.go -package=mocks . Client

import (
	"log"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/config"
)

type Client interface {
	Start() (tgBotAPI.UpdatesChannel, error)
	Send(msg tgBotAPI.MessageConfig) error
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

func (c *client) Start() (tgBotAPI.UpdatesChannel, error) {
	log.Printf("Authorized on account %s", c.tgBot.Self.UserName)
	return c.initUpdatesChannel(), nil
}

func (c *client) Send(msg tgBotAPI.MessageConfig) error {
	_, err := c.tgBot.Send(msg)
	return err
}

func (c *client) initUpdatesChannel() tgBotAPI.UpdatesChannel {
	u := tgBotAPI.NewUpdate(c.cfg.Offset())
	u.Timeout = c.cfg.Timeout()

	return c.tgBot.GetUpdatesChan(u)
}
