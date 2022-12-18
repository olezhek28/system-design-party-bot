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
		timezone = time.Duration(user[0].Timezone.Int64) * time.Hour
	}

	meets, err := s.meetingRepository.GetMeetingsByStatus(ctx, model.MeetingStatusNew)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	meets = excludeDuplicateMeetings(meets)

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
			StartDate:                m.StartDate.Add(timezone).Format(timeFormat),
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

// TODO refactor, del n*n complexity
func excludeDuplicateMeetings(meetings []*model.Meeting) []*model.Meeting {
	res := make([]*model.Meeting, 0, len(meetings))
	for i := 0; i < len(meetings); i++ {
		isDuplicate := false
		for j := i + 1; j < len(meetings); j++ {
			if meetings[i].SpeakerID == meetings[j].ListenerID &&
				meetings[i].ListenerID == meetings[j].SpeakerID &&
				meetings[i].TopicID == meetings[j].TopicID &&
				meetings[i].StartDate == meetings[j].StartDate {
				isDuplicate = true
				break
			}
		}

		if !isDuplicate {
			res = append(res, meetings[i])
		}
	}

	return res
}
