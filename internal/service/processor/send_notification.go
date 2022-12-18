package processor

import (
	"time"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/template"
)

func (s *Service) sendNotification(initiator *model.Student, topicName string, startDate time.Time, recipientID int64) error {
	if initiator == nil {
		return nil
	}

	t, err := helper.ExecuteTemplate(template.NotificationAfterCreate, struct {
		FirstName        string
		LastName         string
		TelegramUsername string
		StartDate        string
		Emoji            string
		TopicName        string
	}{
		FirstName:        initiator.FirstName,
		LastName:         initiator.LastName,
		TelegramUsername: initiator.TelegramUsername,
		StartDate:        startDate.Format(timeFormat),
		Emoji:            model.GetEmoji(model.FoodEmojis),
		TopicName:        topicName,
	})
	if err != nil {
		return err
	}

	reply := tgBotAPI.NewMessage(recipientID, t)
	reply.ParseMode = tgBotAPI.ModeHTML

	err = s.telegramClient.Send(reply)
	if err != nil {
		return err
	}

	return nil
}
