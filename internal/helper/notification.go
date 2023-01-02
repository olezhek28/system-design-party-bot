package helper

import (
	"time"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/model"
)

func GetNotification(initiator *model.Student, unitName string, topicName string, startDate time.Time, recipientID int64, tmpl string) (tgBotAPI.MessageConfig, error) {
	if initiator == nil {
		return tgBotAPI.MessageConfig{}, nil
	}

	t, err := ExecuteTemplate(tmpl, struct {
		FirstName        string
		LastName         string
		TelegramUsername string
		StartDate        string
		Emoji            string
		TopicName        string
		UnitName         string
	}{
		FirstName:        initiator.FirstName,
		LastName:         initiator.LastName,
		TelegramUsername: initiator.TelegramUsername,
		StartDate:        startDate.Format(model.TimeFormat),
		Emoji:            model.GetEmoji(model.FoodEmojis),
		TopicName:        topicName,
		UnitName:         unitName,
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, nil
	}

	reply := tgBotAPI.NewMessage(recipientID, t)
	reply.ParseMode = tgBotAPI.ModeHTML

	return reply, nil
}
