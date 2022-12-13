package processor

import (
	"context"
	"fmt"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/converter"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/pkg/http/telegram"
	meetingRepository "github.com/olezhek28/system-design-party-bot/internal/repository/meeting"
	topicRepository "github.com/olezhek28/system-design-party-bot/internal/repository/topic"
)

type Handler func(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error)

type Service struct {
	telegramClient telegram.Client

	meetingRepository meetingRepository.Repository
	topicRepository   topicRepository.Repository
}

func NewService(
	telegramClient telegram.Client,
	meetingRepository meetingRepository.Repository,
	topicRepository topicRepository.Repository,
) *Service {
	return &Service{
		telegramClient:    telegramClient,
		meetingRepository: meetingRepository,
		topicRepository:   topicRepository,
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

		err = s.telegramClient.Send(reply)
		if err != nil {
			fmt.Printf("failed to send message, err: %s\n", err.Error())
		}
	}

	return nil
}

func (s *Service) getCommandMap() map[string]Handler {
	return map[string]Handler{
		"start":        s.Start,
		"find_speaker": s.FindSpeaker,
		"list_topics":  s.ListTopics,
	}
}

func getCommandKeyboard() tgBotAPI.ReplyKeyboardMarkup {
	return tgBotAPI.NewReplyKeyboard(
		tgBotAPI.NewKeyboardButtonRow(
			tgBotAPI.NewKeyboardButton("/list_topics"),
			tgBotAPI.NewKeyboardButton("2"),
			tgBotAPI.NewKeyboardButton("3"),
		),
		tgBotAPI.NewKeyboardButtonRow(
			tgBotAPI.NewKeyboardButton("4"),
			tgBotAPI.NewKeyboardButton("5"),
			tgBotAPI.NewKeyboardButton("6"),
		),
	)
}
