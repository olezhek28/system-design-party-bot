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
	isExist, err := s.studentRepository.IsExist(ctx, msg.From.ID)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	if !isExist {
		err = s.studentRepository.Create(ctx, &model.Student{
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
	reply.ParseMode = tgBotAPI.ModeHTML

	return reply, nil
}

func getStartKeyboard() tgBotAPI.ReplyKeyboardMarkup {
	return tgBotAPI.NewReplyKeyboard(
		tgBotAPI.NewKeyboardButtonRow(
			tgBotAPI.NewKeyboardButton(fmt.Sprintf("/%s", command.CreateMeeting)),
			tgBotAPI.NewKeyboardButton(fmt.Sprintf("/%s", command.GetStudents)),
		),
		tgBotAPI.NewKeyboardButtonRow(
			tgBotAPI.NewKeyboardButton(fmt.Sprintf("/%s", command.GetSocialConnections)),
			tgBotAPI.NewKeyboardButton(fmt.Sprintf("/%s", command.Help)),
		),
	)
}
