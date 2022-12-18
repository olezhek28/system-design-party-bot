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

	if len(msg.Arguments) == 0 {
		reply := tgBotAPI.NewMessage(msg.From.ID, "Сам выберешь партнёра или мне подыскать наилучший варик?\n")
		reply.ReplyMarkup = getChoiceModeKeyboard()
		return reply, nil
	}
	if len(msg.Arguments) < 3 {
		return tgBotAPI.MessageConfig{}, errors.New("not enough arguments")
	}
	if len(msg.Arguments) < 8 {
		var topicID int64
		topicID, err = strconv.ParseInt(msg.Arguments[0], 10, 64)
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}

		var speakerID int64
		speakerID, err = strconv.ParseInt(msg.Arguments[1], 10, 64)
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}

		var listenerID int64
		listenerID, err = strconv.ParseInt(msg.Arguments[2], 10, 64)
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}

		if speakerID == 0 {
			speakerID, err = s.getBestSpeaker(ctx, topicID, listenerID)
			if err != nil {
				return tgBotAPI.MessageConfig{}, err
			}
		}

		var speakersInfo []*model.Student
		speakersInfo, err = s.studentRepository.GetStudentByIDs(ctx, []int64{speakerID})
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}
		if len(speakersInfo) == 0 {
			return tgBotAPI.MessageConfig{}, errors.New("speaker not found")
		}

		var topicsInfo []*model.Topic
		topicsInfo, err = s.topicRepository.GetTopicsByIDs(ctx, []int64{topicID})
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}
		if len(topicsInfo) == 0 {
			return tgBotAPI.MessageConfig{}, errors.New("topic not found")
		}

		var count int64
		count, err = s.meetingRepository.GetSpeakerCountByTopic(ctx, topicID, speakerID)
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}

		var t string
		t, err = helper.ExecuteTemplate(template.CreateMeetingPrepare, struct {
			FirstName        string
			LastName         string
			TelegramUsername string
			Emoji            string
			TopicName        string
			Count            int64
		}{
			FirstName:        speakersInfo[0].FirstName,
			LastName:         speakersInfo[0].LastName,
			TelegramUsername: speakersInfo[0].TelegramUsername,
			Emoji:            model.GetEmoji(model.FoodEmojis),
			TopicName:        topicsInfo[0].Name,
			Count:            count,
		})
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}

		reply := tgBotAPI.NewMessage(msg.From.ID, t)
		reply.ReplyMarkup = getPickMonthKeyboard(msg.Arguments)
		reply.ParseMode = tgBotAPI.ModeHTML

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

	if speakerID == 0 {
		speakerID, err = s.getBestSpeaker(ctx, topicID, listenerID)
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}
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

	topic, err := s.topicRepository.GetTopicsByIDs(ctx, []int64{topicID})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}
	if len(topic) == 0 {
		return tgBotAPI.MessageConfig{}, errors.New("topic not found")
	}

	count, err := s.meetingRepository.GetSpeakerCountByTopic(ctx, topicID, speakerID)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	t, err := helper.ExecuteTemplate(template.CreateMeetingDescription, struct {
		FirstName        string
		LastName         string
		TelegramUsername string
		StartDate        string
		Emoji            string
		TopicName        string
		Count            int64
	}{
		FirstName:        speakersInfo[0].FirstName,
		LastName:         speakersInfo[0].LastName,
		TelegramUsername: speakersInfo[0].TelegramUsername,
		StartDate:        startDateLocal.Format(model.TimeFormat),
		Emoji:            model.GetEmoji(model.FoodEmojis),
		TopicName:        topic[0].Name,
		Count:            count,
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	listeners, err := s.studentRepository.GetStudentByIDs(ctx, []int64{listenerID})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}
	if len(listeners) == 0 {
		return tgBotAPI.MessageConfig{}, errors.New("speaker not found")
	}

	startDateListener := startDateUTC
	if speakersInfo[0].Timezone.Valid {
		startDateListener = startDateListener.Add(time.Duration(speakersInfo[0].Timezone.Int64) * time.Hour)
	}

	notificationMsg, err := helper.GetNotification(listeners[0], topic[0].Name, startDateListener, speakersInfo[0].TelegramID, template.NotificationAfterCreate)
	if err != nil {
		fmt.Printf("error while getting notification message: %v\n", err)
	}

	err = s.telegramClient.Send(notificationMsg)
	if err != nil {
		fmt.Printf("error while sending notification: %v\n", err)
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, t)
	reply.ReplyMarkup = getMeetingKeyboard(meetingID1, meetingID2)
	reply.ParseMode = tgBotAPI.ModeHTML

	return reply, nil
}

func (s *Service) getBestSpeaker(ctx context.Context, topicID int64, listenerID int64) (int64, error) {
	stats, err := s.meetingRepository.GetSpeakersStats(ctx, topicID, listenerID)
	if err != nil {
		return 0, err
	}

	speakerID := helper.GetInexperiencedSpeaker(stats)
	if speakerID == 0 {
		var speakerInfo *model.Student
		speakerInfo, err = s.studentRepository.GetRandomStudent(ctx, listenerID)
		if err != nil {
			return 0, err
		}
		if speakerInfo == nil {
			return 0, errors.New("no speakers")
		}
		if speakerInfo.ID == listenerID {
			return 0, errors.New("🚫 Что-то кроме тебя я никого пока не знаю:( Зови друзей сюда и начнём движение.")
		}

		speakerID = speakerInfo.ID
	}

	return speakerID, nil
}

func getChoiceModeKeyboard() tgBotAPI.InlineKeyboardMarkup {
	return tgBotAPI.NewInlineKeyboardMarkup(
		tgBotAPI.NewInlineKeyboardRow(
			tgBotAPI.NewInlineKeyboardButtonData(
				fmt.Sprintf("%s Выберу сам", model.GetEmoji(model.TransportEmoji)),
				fmt.Sprintf("/%s %t", command.GetStudents, false),
			),
			tgBotAPI.NewInlineKeyboardButtonData(
				fmt.Sprintf("%s Найди плииииз", model.GetEmoji(model.ThingsEmojis)),
				fmt.Sprintf("/%s %t", command.ListTopics, true),
			),
		),
	)
}

func getPickMonthKeyboard(args []string) tgBotAPI.InlineKeyboardMarkup {
	now := time.Now()

	yearCurrentMonth := now.Year()
	currentMonth := int64(now.Month())

	yearNextMonth := now.Year()
	nextMonth := currentMonth + 1

	if nextMonth > 12 {
		nextMonth = 1
		yearNextMonth++
	}

	buttonsInfo := []*model.TelegramButtonInfo{
		{
			Text: helper.GetMonthList()[currentMonth],
			Data: fmt.Sprintf("/%s %d %d %s", command.PickMonth, currentMonth, yearCurrentMonth, helper.SliceToString(args)),
		},
		{
			Text: helper.GetMonthList()[nextMonth],
			Data: fmt.Sprintf("/%s %d %d %s", command.PickMonth, nextMonth, yearNextMonth, helper.SliceToString(args)),
		},
	}

	return helper.BuildKeyboard(buttonsInfo, 2)
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
