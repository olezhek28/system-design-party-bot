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

func (s *Service) PickDay(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	if len(msg.Arguments) < 6 {
		return tgBotAPI.MessageConfig{}, errors.New("no arguments")
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, "Выбери час\n")
	reply.ReplyMarkup = getPickHourKeyboard(msg.Arguments)

	return reply, nil
}

func getPickHourKeyboard(args []string) tgBotAPI.InlineKeyboardMarkup {
	var buttonsInfo []*model.TelegramButtonInfo
	for _, h := range helper.GetHours() {
		buttonsInfo = append(buttonsInfo, &model.TelegramButtonInfo{
			Text: fmt.Sprintf("%d", h),
			Data: fmt.Sprintf("/%s %d %s", command.PickHour, h, helper.SliceToString(args)),
		})
	}

	return helper.BuildKeyboard(buttonsInfo)
}
