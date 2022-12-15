package processor

import (
	"context"
	"fmt"
	"strconv"
	"time"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/model/command"
	"github.com/olezhek28/system-design-party-bot/internal/template"
	"github.com/pkg/errors"
)

func (s *Service) CreateMeeting(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	if len(msg.Arguments) < 3 {
		return tgBotAPI.MessageConfig{}, errors.New("not enough arguments")
	}

	topicID, err := strconv.ParseInt(msg.Arguments[0], 10, 64)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	speakerID, err := strconv.ParseInt(msg.Arguments[1], 10, 64)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	listenerID, err := strconv.ParseInt(msg.Arguments[2], 10, 64)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	meetingID1, err := s.meetingRepository.CreateMeeting(ctx, &model.Meeting{
		TopicID:    topicID,
		Status:     model.MeetingStatusNew,
		StartDate:  time.Now(),
		SpeakerID:  speakerID,
		ListenerID: listenerID,
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	meetingID2, err := s.meetingRepository.CreateMeeting(ctx, &model.Meeting{
		TopicID:    topicID,
		Status:     model.MeetingStatusNew,
		StartDate:  time.Now(),
		SpeakerID:  listenerID,
		ListenerID: speakerID,
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	speakersInfo, err := s.studentRepository.GetStudentByIDs(ctx, []int64{speakerID})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}
	if len(speakersInfo) == 0 {
		return tgBotAPI.MessageConfig{}, errors.New("speaker not found")
	}

	t, err := helper.ExecuteTemplate(template.CreateMeetingDescription, struct {
		FirstName string
		LastName  string
	}{
		FirstName: speakersInfo[0].FirstName,
		LastName:  speakersInfo[0].LastName,
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, t)
	reply.ReplyMarkup = getMeetingKeyboard(meetingID1, meetingID2)

	return reply, nil
}

func getMeetingKeyboard(meetingID1 int64, meetingID2 int64) tgBotAPI.InlineKeyboardMarkup {
	return tgBotAPI.NewInlineKeyboardMarkup(
		tgBotAPI.NewInlineKeyboardRow(
			tgBotAPI.NewInlineKeyboardButtonData(
				"Встреча состоялась",
				fmt.Sprintf("/%s %d %d", command.FinishMeeting, meetingID1, meetingID2),
			),
			tgBotAPI.NewInlineKeyboardButtonData(
				"Встреча отменена",
				fmt.Sprintf("/%s %d %d", command.CancelMeeting, meetingID1, meetingID2),
			),
		),
	)
}
