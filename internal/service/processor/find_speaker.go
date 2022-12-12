package processor

import (
	"fmt"

	"github.com/olezhek28/system-design-party-bot/internal/model"
)

func (s *Service) FindSpeaker(msg *model.TelegramMessage) (string, error) {
	return fmt.Sprintf("Hello, %s %s. Your command=%s, args=%s",
		msg.From.FirstName, msg.From.LastName, msg.Command, msg.Arguments), nil
}
