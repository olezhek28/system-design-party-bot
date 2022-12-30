package processor

import (
	"context"
	"strconv"
	"strings"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/model/command"
	"github.com/pkg/errors"
)

func (s *Service) SetTimezone(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	if len(msg.Arguments) == 0 {
		return tgBotAPI.MessageConfig{}, errors.New("no arguments")
	}

	tz := strings.Split(msg.Arguments[0], ":")
	if len(tz) < 1 {
		return tgBotAPI.MessageConfig{}, errors.New("invalid timezone")
	}

	hours, err := strconv.Atoi(tz[0])
	if err != nil {
		return tgBotAPI.MessageConfig{}, errors.New("invalid timezone")
	}

	minutes := 0
	if len(tz) == 2 {
		minutes, err = strconv.Atoi(tz[1])
		if err != nil {
			return tgBotAPI.MessageConfig{}, errors.New("invalid timezone")
		}
	}

	if hours < -12 || hours > 12 {
		return tgBotAPI.MessageConfig{}, errors.New("invalid timezone")
	}
	if minutes < 0 || minutes > 59 {
		return tgBotAPI.MessageConfig{}, errors.New("invalid timezone")
	}

	timezone := int64(hours*60 + minutes)

	res, err := s.studentRepository.SetTimezone(ctx, msg.From.ID, timezone)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}
	if res == 0 {
		return tgBotAPI.NewMessage(msg.From.ID, "Кажется ты не зарегистрирован:( Для этого нажми /"+command.Start), nil
	}

	return tgBotAPI.NewMessage(msg.From.ID, "Временная зона успешно установлена\n"), nil
}
