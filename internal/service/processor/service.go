package processor

import (
	"context"
	"fmt"
	"strings"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/converter"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/model/command"
	"github.com/olezhek28/system-design-party-bot/internal/pkg/http/telegram"
	meetingRepository "github.com/olezhek28/system-design-party-bot/internal/repository/meeting"
	studentRepository "github.com/olezhek28/system-design-party-bot/internal/repository/student"
	topicRepository "github.com/olezhek28/system-design-party-bot/internal/repository/topic"
	"github.com/pkg/errors"
)

type Handler func(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error)

type Service struct {
	telegramClient telegram.Client

	meetingRepository meetingRepository.Repository
	topicRepository   topicRepository.Repository
	studentRepository studentRepository.Repository
}

func NewService(
	telegramClient telegram.Client,
	meetingRepository meetingRepository.Repository,
	topicRepository topicRepository.Repository,
	studentRepository studentRepository.Repository,
) *Service {
	return &Service{
		telegramClient:    telegramClient,
		meetingRepository: meetingRepository,
		topicRepository:   topicRepository,
		studentRepository: studentRepository,
	}
}

func (s *Service) Run(ctx context.Context) error {
	for event := range s.telegramClient.Start() {
		var err error
		var reply tgBotAPI.MessageConfig
		if event.Message != nil {
			msg := converter.MessageToTelegramMessage(event.Message)
			if msg == nil {
				reply = tgBotAPI.NewMessage(event.Message.Chat.ID, "failed to convert message")
			}

			reply, err = s.executeCommand(ctx, msg)
			if err != nil {
				reply = tgBotAPI.NewMessage(event.Message.Chat.ID, err.Error())
			}
		} else {
			reply, err = s.executeCallback(ctx, event)
			if err != nil {
				reply = tgBotAPI.NewMessage(event.CallbackQuery.Message.Chat.ID, err.Error())
			}
		}

		err = s.telegramClient.Send(reply)
		if err != nil {
			fmt.Printf("failed to send message, err: %s\n", err.Error())
		}
	}

	return nil
}

func (s *Service) executeCommand(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	handler, ok := s.getCommandMap()[msg.Command]
	if !ok {
		return tgBotAPI.MessageConfig{}, errors.New("unknown command")
	}

	reply, err := handler(ctx, msg)
	if err != nil {
		return tgBotAPI.MessageConfig{}, errors.Wrap(err, "failed to execute command")
	}

	return reply, nil
}

func (s *Service) executeCallback(ctx context.Context, event tgBotAPI.Update) (tgBotAPI.MessageConfig, error) {
	if strings.HasPrefix(event.CallbackQuery.Data, "/") {
		msg := converter.CallbackDataToTelegramMessage(event.CallbackQuery)
		if msg == nil {
			return tgBotAPI.MessageConfig{}, errors.New("failed to convert message")
		}

		reply, err := s.executeCommand(ctx, msg)
		if err != nil {
			return tgBotAPI.MessageConfig{}, errors.Wrap(err, "failed to execute command")
		}

		reply.BaseChat.ChatID = event.CallbackQuery.Message.Chat.ID
		return reply, nil
	}

	callback := tgBotAPI.NewCallback(event.CallbackQuery.ID, event.CallbackQuery.Data)
	_, err := s.telegramClient.Request(callback)
	if err != nil {
		return tgBotAPI.MessageConfig{}, errors.Wrap(err, "failed to request callback")
	}

	return tgBotAPI.NewMessage(event.CallbackQuery.Message.Chat.ID, event.CallbackQuery.Data), nil
}

// TODO добавить для каждой команды свой обработчик аргументов в одном месте, а не в каждом обработчике
func (s *Service) getCommandMap() map[string]Handler {
	return map[string]Handler{
		command.Start:             s.Start,
		command.FindSpeaker:       s.FindSpeaker,
		command.ListTopics:        s.ListTopics,
		command.GetStatsBySpeaker: s.GetStatsBySpeaker,
		command.GetTopicStats:     s.GetTopicStats,
		command.CreateMeeting:     s.CreateMeeting,
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
