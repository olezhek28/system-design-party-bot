package processor

import (
	"context"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/template"
)

func (s *Service) Help(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	t, err := helper.ExecuteTemplate(template.HelpMsg, nil)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, t)

	return reply, nil
}

// TODO на воскресенье
// 1. Сделать нотификацию о встрече (по времени) - сперва при создании уведомление обоим присылать, потом за 10 минут до начала
