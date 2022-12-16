package processor

import (
	"context"
	"strconv"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/model/command"
	"github.com/pkg/errors"
)

func (s *Service) SetTimezone(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	if len(msg.Arguments) == 0 {
		return tgBotAPI.MessageConfig{}, errors.New("no arguments")
	}

	timezone, err := strconv.ParseInt(msg.Arguments[0], 10, 64)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	res, err := s.studentRepository.SetTimezone(ctx, msg.From.ID, timezone)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}
	if res == 0 {
		return tgBotAPI.NewMessage(msg.From.ID, "Кажется ты не зарегистрирован:( Для этого нажми /"+command.Start), nil
	}

	return tgBotAPI.NewMessage(msg.From.ID, "Временная зона успешно установлена\n"), nil
}
