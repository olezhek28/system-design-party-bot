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
	if event.CallbackQuery == nil {
		return tgBotAPI.MessageConfig{}, nil
	}

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
		command.Start:                   s.Start,
		command.ListTopics:              s.ListTopics,
		command.GetStatsBySpeaker:       s.GetStatsBySpeaker,
		command.GetTopicStats:           s.GetTopicStats,
		command.CreateMeeting:           s.CreateMeeting,
		command.FinishMeeting:           s.FinishMeeting,
		command.CancelMeeting:           s.CancelMeeting,
		command.GetStudents:             s.GetStudents,
		command.GetCalendar:             s.GetCalendar,
		command.GetAllCalendar:          s.GetAllCalendar,
		command.SetTimezone:             s.SetTimezone,
		command.GetTimezone:             s.GetTimezone,
		command.GetSocialConnections:    s.GetSocialConnections,
		command.GetAllSocialConnections: s.GetAllSocialConnections,

		command.PickMonth: s.PickMonth,
		command.PickDay:   s.PickDay,
		command.PickHour:  s.PickHour,
		command.PickMin:   s.PickMin,

		command.Help: s.Help,

		// TODO создание гугл митс встречь
		// TODO присылание ссылки на гугл митс встреч с датой и временем в общий чат
		// TODO добавить команду /add_topic
		// TODO добавить крон, который чекает встречи по времени и если осталось чутка до встречи напоминает о ней участникам
		// TODO добавить в шаблоны хтмл тегов
		// TODO Хранить инфу о ДР, чтоб поздравлять

		// !!!! TODO Сделать команду, которая позволяет какое-то сообщение рассылать всем участникам

		// TODO ПЕРВОСТЕПЕННО
		// 1. Добавить кнопки управления встречей в календарь каждого участника
		// 2. Доработать механизм выбора участника на рассказ
		// 3. Добавить выбор категории тем, чтобы не только систем дизайн мог быть
		// 4. В социальных связях не показывать самого себя
	}
}

// docker exec -t system-design-party-bot_db_1 pg_dumpall -c -U system-design-party-bot-user > dump_$(date +%Y-%m-%d_%H_%M_%S).sql
// docker exec -t system-design-party-bot_db_1 pg_dumpall -c -U system-design-party-bot-user | gzip > ./dump_$(date +"%Y-%m-%d_%H_%M_%S").gz
// cat dump_2022-12-19_12_50_46.sql | docker exec -i system-design-party-bot_db_1 psql -U system-design-party-bot-user -d system-design-party-bot
// gunzip < dump_2022-12-19_13_57_00.gz | docker exec -i system-design-party-bot_db_1 psql -U system-design-party-bot-user -d system-design-party-bot
