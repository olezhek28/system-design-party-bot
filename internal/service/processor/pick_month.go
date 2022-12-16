package processor

import (
	"context"
	"fmt"
	"strconv"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/model/command"
	"github.com/pkg/errors"
)

func (s *Service) PickMonth(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	if len(msg.Arguments) < 5 {
		return tgBotAPI.MessageConfig{}, errors.New("no arguments")
	}

	month, err := strconv.ParseInt(msg.Arguments[0], 10, 64)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	year, err := strconv.ParseInt(msg.Arguments[1], 10, 64)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, "Выбери день\n")
	reply.ReplyMarkup = getPickDaysKeyboard(helper.GetDaysInMonth(year, month), msg.Arguments)

	return reply, nil
}

func getPickDaysKeyboard(dayList []int64, args []string) tgBotAPI.InlineKeyboardMarkup {
	var buttonsInfo []*model.TelegramButtonInfo
	for _, d := range dayList {
		buttonsInfo = append(buttonsInfo, &model.TelegramButtonInfo{
			Text: fmt.Sprintf("%d", d),
			Data: fmt.Sprintf("/%s %d %s", command.PickDay, d, helper.SliceToString(args)),
		})
	}

	return helper.BuildKeyboard(buttonsInfo)
}
