package processor

import (
	"context"
	"fmt"
	"strconv"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/model"
)

func (s *Service) CreateMeeting(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	if len(msg.Arguments) < 2 {
		return tgBotAPI.MessageConfig{}, fmt.Errorf("not enough arguments")
	}

	speakerID, err := strconv.ParseInt(msg.Arguments[0], 10, 64)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	listenerID, err := strconv.ParseInt(msg.Arguments[1], 10, 64)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, fmt.Sprintf("Meeting %d and %d created", speakerID, listenerID))

	return reply, nil
}
