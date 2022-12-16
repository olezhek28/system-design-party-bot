package processor

import (
	"context"
	"strings"
	"time"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/model/command"
	"github.com/olezhek28/system-design-party-bot/internal/template"
	"github.com/pkg/errors"
)

func (s *Service) GetCalendar(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	user, err := s.studentRepository.GetStudentByTelegramChatIDs(ctx, []int64{msg.From.ID})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}
	if len(user) == 0 {
		return tgBotAPI.NewMessage(msg.From.ID, "Кажется ты не зарегистрирован:( Для этого нажми /"+command.Start), nil
	}

	var timezone time.Duration
	if user[0].Timezone.Valid {
		timezone = time.Duration(user[0].Timezone.Int64) * time.Hour
	}

	meets, err := s.meetingRepository.GetMeetingsByStatus(ctx, model.MeetingStatusNew)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	topicMap, err := s.getTopicMap(ctx, meets)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	speakerMap, err := s.getSpeakerMap(ctx, meets)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	res := strings.Builder{}
	t, err := helper.ExecuteTemplate(template.CalendarDescription, nil)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	res.WriteString(t + "\n")

	for _, m := range meets {
		topic, ok := topicMap[m.TopicID]
		if !ok {
			errors.Errorf("topic with id %d not found\n", m.TopicID)
			continue
		}

		speaker, ok := speakerMap[m.SpeakerID]
		if !ok {
			errors.Errorf("speaker with id %d not found\n", m.SpeakerID)
			continue
		}

		listener, ok := speakerMap[m.ListenerID]
		if !ok {
			errors.Errorf("listener with id %d not found\n", m.ListenerID)
			continue
		}

		t, err = helper.ExecuteTemplate(template.CalendarMeeting, struct {
			SpeakerFirstName         string
			SpeakerLastName          string
			SpeakerTelegramUsername  string
			ListenerFirstName        string
			ListenerLastName         string
			ListenerTelegramUsername string
			TopicName                string
			StartDate                string
		}{
			SpeakerFirstName:         speaker.FirstName,
			SpeakerLastName:          speaker.LastName,
			SpeakerTelegramUsername:  speaker.TelegramUsername,
			ListenerFirstName:        listener.FirstName,
			ListenerLastName:         listener.LastName,
			ListenerTelegramUsername: listener.TelegramUsername,
			TopicName:                topic.Name,
			StartDate:                m.StartDate.Add(timezone).Format("02-Jan-2006 15:04"),
		})
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}

		res.WriteString(t + "\n")
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, res.String())
	return reply, nil
}
