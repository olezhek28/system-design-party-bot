package converter

import (
	"strings"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/model"
)

func ToUser(user *tgBotAPI.User) *model.User {
	if user == nil {
		return nil
	}

	return &model.User{
		ID:        user.ID,
		IsBot:     user.IsBot,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
	}
}

func ToTelegramMessage(msg *tgBotAPI.Message) *model.TelegramMessage {
	if msg == nil {
		return nil
	}

	return &model.TelegramMessage{
		ID:        msg.MessageID,
		From:      ToUser(msg.From),
		Text:      msg.Text,
		Command:   msg.Command(),
		Arguments: strings.Fields(msg.CommandArguments()),
	}
}
