package processor

import (
	"context"
	"fmt"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/model/command"
)

func (s *Service) GetTimezone(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	timezone, err := s.studentRepository.GetTimezone(ctx, msg.From.ID)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}
	if timezone == -1 {
		return tgBotAPI.NewMessage(msg.From.ID, "Кажется ты не зарегистрирован:( Для этого нажми /"+command.Start), nil
	}

	var mark string
	if timezone > 0 {
		mark = "+"
	}
	reply := fmt.Sprintf("Твоя временная зона: UTC%s%d\n", mark, timezone)

	return tgBotAPI.NewMessage(msg.From.ID, reply), nil
}