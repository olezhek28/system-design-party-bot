package scheduler

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/template"
)

func (s *Service) sendNotification(ctx context.Context) error {
	now := time.Now().UTC()

	period, err := strconv.ParseInt(model.ReminderTime, 10, 64)
	if err != nil {
		return err
	}

	begin := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, now.Location()).Add(time.Duration(period) * time.Minute)

	meets, err := s.meetingRepository.GetMeetingsByRange(ctx, begin)
	if err != nil {
		return err
	}

	meets = helper.ExcludeDuplicateMeetings(meets)

	fmt.Printf("i found %d meetings by %v\n", len(meets), begin)

	for _, meet := range meets {
		speaker, errMeet := s.studentRepository.GetStudentByIDs(ctx, []int64{meet.SpeakerID})
		if errMeet != nil {
			return errMeet
		}
		if len(speaker) == 0 {
			return fmt.Errorf("speaker with id %d not found", meet.SpeakerID)
		}

		listener, errMeet := s.studentRepository.GetStudentByIDs(ctx, []int64{meet.ListenerID})
		if errMeet != nil {
			return errMeet
		}
		if len(listener) == 0 {
			return fmt.Errorf("listener with id %d not found", meet.ListenerID)
		}

		topic, errMeet := s.topicRepository.GetTopicsByIDs(ctx, []int64{meet.TopicID})
		if errMeet != nil {
			return errMeet
		}
		if len(topic) == 0 {
			return fmt.Errorf("topic with id %d not found", meet.TopicID)
		}

		err = s.sendNotificationToPartners(speaker[0], listener[0], topic[0], meet.StartDate)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) sendNotificationToPartners(speaker *model.Student, listener *model.Student, topic *model.Topic, startDate time.Time) error {
	speakerMsg, err := helper.GetNotification(listener, topic.Name, startDate, speaker.TelegramID, template.NotificationBeforeStart)
	if err != nil {
		fmt.Printf("error while getting notification message: %v\n", err)
		return err
	}

	listenerMsg, err := helper.GetNotification(speaker, topic.Name, startDate, listener.TelegramID, template.NotificationBeforeStart)
	if err != nil {
		fmt.Printf("error while getting notification message: %v\n", err)
		return err
	}

	err = s.telegramClient.Send(speakerMsg)
	if err != nil {
		fmt.Printf("error while sending notification to speaker: %v\n", err)
		return err
	}

	err = s.telegramClient.Send(listenerMsg)
	if err != nil {
		fmt.Printf("error while sending notification to listener: %v\n", err)
		return err
	}

	return nil
}
