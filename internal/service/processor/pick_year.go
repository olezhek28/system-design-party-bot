package processor

import (
	"context"
	"fmt"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/model/command"
	"github.com/pkg/errors"
)

func (s *Service) PickYear(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	if len(msg.Arguments) < 4 {
		return tgBotAPI.MessageConfig{}, errors.New("no arguments")
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, "Выбери месяц\n")
	reply.ReplyMarkup = getPickMonthKeyboard(msg.Arguments)

	return reply, nil
}

// TODO: группировать дни как календарь с днями недели
// TODO: добавить кнопку "Сегодня"
// TODO: добавить кнопку "Завтра"
// TODO: добавить кнопку "Послезавтра"
// TODO: выдавать текущий месяц и следующий
func getPickMonthKeyboard(args []string) tgBotAPI.InlineKeyboardMarkup {
	var buttonsInfo []*model.TelegramButtonInfo
	for i, m := range helper.GetMonthList() {
		buttonsInfo = append(buttonsInfo, &model.TelegramButtonInfo{
			Text: m,
			Data: fmt.Sprintf("/%s %d %s", command.PickMonth, i+1, helper.SliceToString(args)),
		})
	}

	return helper.BuildKeyboard(buttonsInfo)
}
