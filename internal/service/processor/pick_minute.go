package processor

import (
	"context"
	"fmt"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/model/command"
	"github.com/olezhek28/system-design-party-bot/internal/template"
	"github.com/pkg/errors"
)

func (s *Service) PickMin(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	if len(msg.Arguments) < 8 {
		return tgBotAPI.MessageConfig{}, errors.New("no arguments")
	}

	startDate, err := getStartDate(msg.Arguments)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	t, err := helper.ExecuteTemplate(template.MeetingConfirmation, struct {
		StartDate string
		Emoji     string
	}{
		StartDate: startDate.Format(timeFormat),
		Emoji:     model.GetEmoji(model.DrinksEmojis),
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, t)
	reply.ReplyMarkup = getMeetingYesNoKeyboard(msg.Arguments)

	return reply, nil
}

func getMeetingYesNoKeyboard(args []string) tgBotAPI.InlineKeyboardMarkup {
	return tgBotAPI.NewInlineKeyboardMarkup(
		tgBotAPI.NewInlineKeyboardRow(
			tgBotAPI.NewInlineKeyboardButtonData(
				"Да",
				fmt.Sprintf("/%s %s", command.CreateMeeting, helper.SliceToString(args)),
			),
		),
	)
}
