package helper

import "github.com/olezhek28/system-design-party-bot/internal/model"

func SplitSlice(data []*model.TelegramButtonInfo, chunkSize int) [][]*model.TelegramButtonInfo {
	var chunks [][]*model.TelegramButtonInfo

	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize

		if end > len(data) {
			end = len(data)
		}

		chunks = append(chunks, data[i:end])
	}

	return chunks
}
