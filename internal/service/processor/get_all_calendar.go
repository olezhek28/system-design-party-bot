package processor

import (
	"context"
	"database/sql"
	"strings"
	"time"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/model/command"
	meetingRepository "github.com/olezhek28/system-design-party-bot/internal/repository/meeting"
	"github.com/olezhek28/system-design-party-bot/internal/template"
	"github.com/pkg/errors"
)

func (s *Service) GetAllCalendar(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	user, err := s.studentRepository.GetStudentByTelegramChatIDs(ctx, []int64{msg.From.ID})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}
	if len(user) == 0 {
		return tgBotAPI.NewMessage(msg.From.ID, "Кажется ты не зарегистрирован:( Для этого нажми /"+command.Start), nil
	}

	var timezone time.Duration
	if user[0].Timezone.Valid {
		hours := user[0].Timezone.Int64 / 60
		minutes := user[0].Timezone.Int64 % 60

		timezone = time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute
	}

	meets, err := s.meetingRepository.GetList(ctx, &meetingRepository.Query{
		QueryFilter: model.QueryFilter{
			AllData: true,
		},
		Status: sql.NullString{String: model.MeetingStatusNew, Valid: true},
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	meets = helper.ExcludeDuplicateMeetings(meets)

	topicMap, err := s.getTopicMap(ctx, meets)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	speakerMap, err := s.getSpeakerMap(ctx, meets)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	listenerMap, err := s.getListenerMap(ctx, meets)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	res := strings.Builder{}
	t, err := helper.ExecuteTemplate(template.CalendarAllDescription, struct {
		Emoji string
	}{
		Emoji: model.GetEmoji(model.CalendarEmojis),
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	res.WriteString(t + "\n")

	for _, m := range meets {
		unit, ok := topicMap[m.UnitID]
		if !ok {
			errors.Errorf("unit with id %d not found\n", m.UnitID)
			continue
		}

		topic, ok := unit[m.TopicID]
		if !ok {
			errors.Errorf("topic with id %d not found\n", m.TopicID)
			continue
		}

		speaker, ok := speakerMap[m.SpeakerID]
		if !ok {
			errors.Errorf("speaker with id %d not found\n", m.SpeakerID)
			continue
		}

		listener, ok := listenerMap[m.ListenerID]
		if !ok {
			errors.Errorf("listener with id %d not found\n", m.ListenerID)
			continue
		}

		t, err = helper.ExecuteTemplate(template.CalendarAllMeeting, struct {
			SpeakerFirstName         string
			SpeakerLastName          string
			SpeakerTelegramUsername  string
			ListenerFirstName        string
			ListenerLastName         string
			ListenerTelegramUsername string
			TopicName                string
			StartDate                string
			Emoji                    string
		}{
			SpeakerFirstName:         speaker.FirstName,
			SpeakerLastName:          speaker.LastName,
			SpeakerTelegramUsername:  speaker.TelegramUsername,
			ListenerFirstName:        listener.FirstName,
			ListenerLastName:         listener.LastName,
			ListenerTelegramUsername: listener.TelegramUsername,
			TopicName:                topic.Name,
			StartDate:                m.StartDate.Add(timezone).Format(model.TimeFormat),
			Emoji:                    model.GetEmoji(model.DrinksEmojis),
		})
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}

		res.WriteString(t + "\n")
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, res.String())
	return reply, nil
}
