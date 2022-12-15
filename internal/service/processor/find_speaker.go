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
	"github.com/pkg/errors"
)

func (s *Service) FindSpeaker(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	if len(msg.Arguments) == 0 {
		return tgBotAPI.MessageConfig{}, errors.New("no arguments")
	}

	topicID, err := strconv.ParseInt(msg.Arguments[0], 10, 64)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	stats, err := s.meetingRepository.GetSpeakersStats(ctx, topicID, msg.From.ID)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	var speaker *model.Stats
	if len(stats) != 0 {
		speaker = stats[0]
		minCount := stats[0].Count
		for _, stat := range stats {
			if stat.Count < minCount {
				minCount = stat.Count
				speaker = stat
			}
		}
	} else {
		var speakerInfo *model.Student
		speakerInfo, err = s.studentRepository.GetRandomStudent(ctx, msg.From.ID)
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}
		if speakerInfo == nil {
			return tgBotAPI.MessageConfig{}, errors.New("no speakers")
		}

		if speakerInfo.TelegramID == msg.From.ID {
			return tgBotAPI.NewMessage(msg.From.ID, "🚫 Что-то кроме тебя я никого пока не знаю:( Зови друзей сюда и начнём движение."), nil
		}

		var topicInfo []*model.Topic
		topicInfo, err = s.topicRepository.GetTopicsByIDs(ctx, []int64{topicID})
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}
		if len(topicInfo) == 0 {
			return tgBotAPI.MessageConfig{}, errors.New("topic not found")
		}

		speaker = &model.Stats{
			SpeakerID:               speakerInfo.ID,
			SpeakerFirstName:        speakerInfo.FirstName,
			SpeakerLastName:         speakerInfo.LastName,
			SpeakerTelegramNickname: speakerInfo.TelegramUsername,
			TopicName:               topicInfo[0].Name,
			Count:                   0,
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

	listener, err := s.studentRepository.GetStudentByTelegramChatIDs(ctx, []int64{msg.From.ID})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}
	if len(listener) == 0 {
		return tgBotAPI.MessageConfig{}, errors.New("listener not found")
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, t)
	reply.ReplyMarkup = getCreateMeetingKeyboard(speaker, listener[0], topicID)

	return reply, nil
}

func getCreateMeetingKeyboard(speaker *model.Stats, listener *model.Student, topicID int64) tgBotAPI.InlineKeyboardMarkup {
	return tgBotAPI.NewInlineKeyboardMarkup(
		tgBotAPI.NewInlineKeyboardRow(
			tgBotAPI.NewInlineKeyboardButtonData(
				fmt.Sprintf("Создать встречу с пользователем %s %s", speaker.SpeakerFirstName, speaker.SpeakerLastName),
				fmt.Sprintf("/%s %d %d %d", command.CreateMeeting, topicID, speaker.SpeakerID, listener.ID),
			),
		),
	)
}
