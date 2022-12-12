package app

import (
	"log"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/config"
	"github.com/olezhek28/system-design-party-bot/internal/pkg/http/telegram"
	"github.com/olezhek28/system-design-party-bot/internal/service/processor"
)

type serviceProvider struct {
	telegramClient telegram.Client

	processorService *processor.Service
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GetTelegramClient() telegram.Client {
	if s.telegramClient == nil {
		cfg, err := config.NewConfig()
		if err != nil {
			log.Fatalf("failed to get telegram config: %s", err.Error())
		}

		bot, err := tgBotAPI.NewBotAPI(cfg.TelegramToken())
		if err != nil {
			log.Fatalf("failed to creating new tg client: %s", err.Error())
		}

		s.telegramClient = telegram.NewClient(bot, cfg)
	}

	return s.telegramClient
}

func (s *serviceProvider) GetProcessorService() *processor.Service {
	if s.processorService == nil {
		s.processorService = processor.NewService(s.GetTelegramClient())
	}

	return s.processorService
}
