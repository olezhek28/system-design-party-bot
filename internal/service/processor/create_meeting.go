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
	user, err := s.studentRepository.GetStudentByTelegramChatIDs(ctx, []int64{msg.From.ID})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}
	if len(user) == 0 {
		return tgBotAPI.NewMessage(msg.From.ID, "Кажется ты не зарегистрирован:( Для этого нажми /"+command.Start), nil
	}

	if len(msg.Arguments) < 3 {
		return tgBotAPI.MessageConfig{}, errors.New("not enough arguments")
	}
	if len(msg.Arguments) < 8 {
		reply := tgBotAPI.NewMessage(msg.From.ID, "Выбери год\n")
		reply.ReplyMarkup = getPickYearKeyboard(msg.Arguments)
		return reply, nil
	}

	startDateLocal, err := getStartDate(msg.Arguments)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	startDateUTC := startDateLocal
	if user[0].Timezone.Valid {
		startDateUTC = startDateUTC.Add((-1) * time.Duration(user[0].Timezone.Int64) * time.Hour)
	}

	topicID, err := strconv.ParseInt(msg.Arguments[5], 10, 64)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	speakerID, err := strconv.ParseInt(msg.Arguments[6], 10, 64)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	listenerID, err := strconv.ParseInt(msg.Arguments[7], 10, 64)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	meetingID1, err := s.meetingRepository.CreateMeeting(ctx, &model.Meeting{
		TopicID:    topicID,
		Status:     model.MeetingStatusNew,
		StartDate:  startDateUTC,
		SpeakerID:  speakerID,
		ListenerID: listenerID,
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	meetingID2, err := s.meetingRepository.CreateMeeting(ctx, &model.Meeting{
		TopicID:    topicID,
		Status:     model.MeetingStatusNew,
		StartDate:  startDateUTC,
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
		StartDate string
		Emoji     string
	}{
		FirstName: speakersInfo[0].FirstName,
		LastName:  speakersInfo[0].LastName,
		StartDate: startDateLocal.Format(timeFormat),
		Emoji:     model.GetEmoji(model.FoodEmojis),
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, t)
	reply.ReplyMarkup = getMeetingKeyboard(meetingID1, meetingID2)

	return reply, nil
}

func getPickYearKeyboard(args []string) tgBotAPI.InlineKeyboardMarkup {
	currenYear := time.Now().Year()
	return tgBotAPI.NewInlineKeyboardMarkup(
		tgBotAPI.NewInlineKeyboardRow(
			tgBotAPI.NewInlineKeyboardButtonData(
				fmt.Sprintf("%d", currenYear),
				fmt.Sprintf("/%s %d %s", command.PickYear, currenYear, helper.SliceToString(args)),
			),
			tgBotAPI.NewInlineKeyboardButtonData(
				fmt.Sprintf("%d", currenYear+1),
				fmt.Sprintf("/%s %d %s", command.PickYear, currenYear+1, helper.SliceToString(args)),
			),
		),
	)
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

func getStartDate(args []string) (time.Time, error) {
	min, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	hour, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	day, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	month, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	year, err := strconv.ParseInt(args[4], 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	timestamp := time.Date(int(year), time.Month(month), int(day), int(hour), int(min), 0, 0, time.UTC)

	return timestamp, nil
}
