package processor

import (
	"context"

	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/template"
)

func (s *Service) Start(ctx context.Context, msg *model.TelegramMessage) (string, error) {
	data := struct {
		FirstName string
	}{
		FirstName: msg.From.FirstName,
	}

	res, err := helper.ExecuteTemplate(template.StartMsg, data)
	if err != nil {
		return "", err
	}

	return res, nil
}
