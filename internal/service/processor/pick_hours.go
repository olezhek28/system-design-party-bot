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

func (s *Service) PickHour(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	if len(msg.Arguments) < 7 {
		return tgBotAPI.MessageConfig{}, errors.New("no arguments")
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, "Выбери минуты\n")
	reply.ReplyMarkup = getPickMinuteKeyboard(msg.Arguments)

	return reply, nil
}

func getPickMinuteKeyboard(args []string) tgBotAPI.InlineKeyboardMarkup {
	var buttonsInfo []*model.TelegramButtonInfo
	for _, m := range helper.GetMinutes() {
		buttonsInfo = append(buttonsInfo, &model.TelegramButtonInfo{
			Text: fmt.Sprintf("%d", m),
			Data: fmt.Sprintf("/%s %d %s", command.PickMin, m, helper.SliceToString(args)),
		})
	}

	return helper.BuildKeyboard(buttonsInfo)
}
