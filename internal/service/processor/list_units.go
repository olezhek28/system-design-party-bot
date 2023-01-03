package processor

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/model/command"
	studentRepository "github.com/olezhek28/system-design-party-bot/internal/repository/student"
	unitRepository "github.com/olezhek28/system-design-party-bot/internal/repository/unit"
	"github.com/olezhek28/system-design-party-bot/internal/template"
)

func (s *Service) ListUnits(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	units, err := s.unitRepository.GetList(ctx, &unitRepository.Query{
		QueryFilter: model.QueryFilter{
			AllData: true,
		},
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}
	if len(units) == 0 {
		return tgBotAPI.NewMessage(msg.From.ID, "🚫 Что-то не нашёл в базе ни одного раздела( Спроси у @olezhek28 в чём проблема."), nil
	}

	res := strings.Builder{}
	for _, unit := range units {
		t, errTmpl := helper.ExecuteTemplate(template.UnitDescription, struct {
			ID          int64
			Name        string
			Description string
		}{
			ID:          unit.ID,
			Name:        unit.Name,
			Description: unit.Description,
		})
		if errTmpl != nil {
			return tgBotAPI.MessageConfig{}, errTmpl
		}

		res.WriteString(t + "\n")
	}

	t, err := helper.ExecuteTemplate(template.UnitGuidelines, struct {
		FirstName string
	}{
		FirstName: msg.From.FirstName,
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	res.WriteString(t)

	reply := tgBotAPI.NewMessage(msg.From.ID, res.String())

	var users []*model.Student
	users, err = s.studentRepository.GetList(ctx, &studentRepository.Query{
		QueryFilter: model.QueryFilter{
			AllData: true,
		},
		TelegramIDs: []int64{msg.From.ID},
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}
	if len(users) == 0 {
		return tgBotAPI.NewMessage(msg.From.ID, "Кажется ты не зарегистрирован:( Для этого нажми /"+command.Start), nil
	}

	listenerID := users[0].ID
	var speakerID int64
	flag := false
	if len(msg.Arguments) > 0 {
		flag, err = strconv.ParseBool(msg.Arguments[0])
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}
	}

	if !flag {
		if len(msg.Arguments) == 2 {
			speakerID, err = strconv.ParseInt(msg.Arguments[1], 10, 64)
			if err != nil {
				return tgBotAPI.MessageConfig{}, err
			}
		}
	}

	reply.ReplyMarkup = getStudentsWithListTopicsKeyboard(units, speakerID, listenerID)

	return reply, nil
}

func getStudentsWithListTopicsKeyboard(units []*model.Unit, speakerID int64, listenerID int64) tgBotAPI.InlineKeyboardMarkup {
	var buttonsInfo []*model.TelegramButtonInfo
	for _, unit := range units {
		buttonsInfo = append(buttonsInfo, &model.TelegramButtonInfo{
			Text: fmt.Sprintf("%d", unit.ID),
			Data: fmt.Sprintf("/%s %d %d %d", command.ListTopics, unit.ID, speakerID, listenerID),
		})
	}

	return helper.BuildKeyboard(buttonsInfo, 3)
}
