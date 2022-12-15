package converter

import (
	"strings"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
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

func MessageToTelegramMessage(msg *tgBotAPI.Message) *model.TelegramMessage {
	if msg == nil {
		return nil
	}

	return &model.TelegramMessage{
		ID:        int64(msg.MessageID),
		From:      ToUser(msg.From),
		Text:      msg.Text,
		Command:   msg.Command(),
		Arguments: strings.Fields(msg.CommandArguments()),
	}
}

func CallbackDataToTelegramMessage(query *tgBotAPI.CallbackQuery) *model.TelegramMessage {
	if query == nil {
		return nil
	}

	return &model.TelegramMessage{
		ID:        query.Message.Chat.ID,
		From:      ToUser(query.From),
		Text:      query.Data,
		Command:   helper.Command(query.Data),
		Arguments: helper.CommandArguments(query.Data),
	}
}
