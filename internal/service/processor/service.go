package processor

import (
	"context"
	"fmt"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/converter"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/pkg/http/telegram"
	meetingRepository "github.com/olezhek28/system-design-party-bot/internal/repository/meeting"
)

type Handler func(ctx context.Context, msg *model.TelegramMessage) (string, error)

type Service struct {
	telegramClient telegram.Client

	meetingRepository meetingRepository.Repository
}

func NewService(
	telegramClient telegram.Client,
	meetingRepository meetingRepository.Repository,
) *Service {
	return &Service{
		telegramClient:    telegramClient,
		meetingRepository: meetingRepository,
	}
}

func (s *Service) Run(ctx context.Context) error {
	msgChan, err := s.telegramClient.Start()
	if err != nil {
		return err
	}

	for event := range msgChan {
		msg := converter.ToTelegramMessage(event.Message)
		if msg == nil {
			fmt.Errorf("failed to convert message, update id: %d", event.UpdateID)
		}

		handler, ok := s.getCommandMap()[msg.Command]
		if !ok {
			s.telegramClient.Send(tgBotAPI.NewMessage(msg.From.ID, "Unknown command"))
			continue
		}

		reply, errHandler := handler(ctx, msg)
		if errHandler != nil {
			s.telegramClient.Send(tgBotAPI.NewMessage(msg.From.ID, fmt.Sprintf("failed to execute command: %s", errHandler.Error())))
			continue
		}

		s.telegramClient.Send(tgBotAPI.NewMessage(msg.From.ID, reply))
	}

	return nil
}

func (s *Service) getCommandMap() map[string]Handler {
	return map[string]Handler{
		"start":        s.Start,
		"find_speaker": s.FindSpeaker,
	}
}
