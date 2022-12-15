package helper

import (
	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/model"
)

const columnSize = 3

func BuildKeyboard(buttonsInfo []*model.TelegramButtonInfo) tgBotAPI.InlineKeyboardMarkup {
	chunks := SplitSlice(buttonsInfo, columnSize)

	var buttons [][]tgBotAPI.InlineKeyboardButton
	for i := range chunks {
		var rows []tgBotAPI.InlineKeyboardButton
		for j := range chunks[i] {
			rows = append(rows, tgBotAPI.NewInlineKeyboardButtonData(chunks[i][j].Text, chunks[i][j].Data))
		}

		buttons = append(buttons, rows)
	}

	return tgBotAPI.NewInlineKeyboardMarkup(buttons...)
}
