package processor

import (
	"context"
	"fmt"
	"strconv"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/model/command"
	"github.com/olezhek28/system-design-party-bot/internal/template"
)

func (s *Service) FindSpeaker(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	if len(msg.Arguments) == 0 {
		return tgBotAPI.MessageConfig{}, fmt.Errorf("no arguments")
	}

	topicID, err := strconv.ParseInt(msg.Arguments[0], 10, 64)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	stats, err := s.meetingRepository.GetSpeakers(ctx, topicID)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}
	if len(stats) == 0 {
		return tgBotAPI.MessageConfig{}, nil
	}

	speaker := stats[0]
	minCount := stats[0].Count
	for _, stat := range stats {
		if stat.Count < minCount {
			minCount = stat.Count
			speaker = stat
		}
	}

	t, err := helper.ExecuteTemplate(template.SpeakerDescription, struct {
		FirstName        string
		LastName         string
		TelegramUsername string
		TopicName        string
		Count            int64
	}{
		FirstName:        speaker.SpeakerFirstName,
		LastName:         speaker.SpeakerLastName,
		TelegramUsername: speaker.SpeakerTelegramNickname,
		TopicName:        speaker.TopicName,
		Count:            speaker.Count,
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, t)
	reply.ReplyMarkup = getMeetKeyboard(msg, speaker)

	return reply, nil
}

func getMeetKeyboard(msg *model.TelegramMessage, speaker *model.Stats) tgBotAPI.InlineKeyboardMarkup {
	return tgBotAPI.NewInlineKeyboardMarkup(
		tgBotAPI.NewInlineKeyboardRow(
			tgBotAPI.NewInlineKeyboardButtonData(
				fmt.Sprintf("Создать встречу с пользователем %s %s", speaker.SpeakerFirstName, speaker.SpeakerLastName),
				fmt.Sprintf("/%s %d %d", command.CreateMeeting, speaker.SpeakerID, msg.From.ID),
			),
		),
	)
}
