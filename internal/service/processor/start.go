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
	data := struct {
		FirstName string
	}{
		FirstName: msg.From.FirstName,
	}

	res, err := helper.ExecuteTemplate(template.StartMsg, data)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	resMsg := tgBotAPI.NewMessage(msg.From.ID, res)
	resMsg.ReplyMarkup = getStartKeyboard()

	return resMsg, nil
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
