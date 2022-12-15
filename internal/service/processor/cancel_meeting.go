package processor

import (
	"context"
	"strconv"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/template"
	"github.com/pkg/errors"
)

// TODO не давать менять статус встречи, если она уже завершена
func (s *Service) CancelMeeting(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	if len(msg.Arguments) < 2 {
		return tgBotAPI.MessageConfig{}, errors.New("no arguments")
	}

	meetingID1, err := strconv.ParseInt(msg.Arguments[0], 10, 64)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	meetingID2, err := strconv.ParseInt(msg.Arguments[1], 10, 64)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	err = s.meetingRepository.UpdateMeetingsStatus(ctx, model.MeetingStatusCanceled, []int64{meetingID1, meetingID2})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	t, err := helper.ExecuteTemplate(template.CancelMeetingDescription, nil)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, t)
	reply.ReplyMarkup = getStartKeyboard()

	return reply, nil
}
