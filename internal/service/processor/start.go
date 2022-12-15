package processor

import (
	"context"
	"fmt"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/model/command"
	"github.com/olezhek28/system-design-party-bot/internal/template"
)

func (s *Service) Start(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	isExist, err := s.studentRepository.IsExistStudent(ctx, msg.From.ID)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	if !isExist {
		err = s.studentRepository.CreateStudent(ctx, &model.Student{
			FirstName:        msg.From.FirstName,
			LastName:         msg.From.LastName,
			TelegramID:       msg.From.ID,
			TelegramUsername: msg.From.UserName,
		})
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}
	}

	// TODO менять текст в зависимости от того, существовал ли студент в базе ранее (можно даже дату регистрации выводить)
	res, err := helper.ExecuteTemplate(template.StartMsg, struct {
		FirstName string
	}{
		FirstName: msg.From.FirstName,
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, res)
	reply.ReplyMarkup = getStartKeyboard()

	return reply, nil
}

func getStartKeyboard() tgBotAPI.InlineKeyboardMarkup {
	return tgBotAPI.NewInlineKeyboardMarkup(
		tgBotAPI.NewInlineKeyboardRow(
			tgBotAPI.NewInlineKeyboardButtonData(
				"Показать список тем",
				fmt.Sprintf("/%s", command.ListTopics),
			),
		),
	)
}
