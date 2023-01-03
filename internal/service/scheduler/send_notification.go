package scheduler

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	meetingRepository "github.com/olezhek28/system-design-party-bot/internal/repository/meeting"
	studentRepository "github.com/olezhek28/system-design-party-bot/internal/repository/student"
	topicRepository "github.com/olezhek28/system-design-party-bot/internal/repository/topic"
	unitRepository "github.com/olezhek28/system-design-party-bot/internal/repository/unit"
	"github.com/olezhek28/system-design-party-bot/internal/template"
)

func (s *Service) sendNotification(ctx context.Context) error {
	now := time.Now().UTC()

	period, err := strconv.ParseInt(model.ReminderTime, 10, 64)
	if err != nil {
		return err
	}

	startDate := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, now.Location()).Add(time.Duration(period) * time.Minute)

	meets, err := s.meetingRepository.GetList(ctx, &meetingRepository.Query{
		QueryFilter: model.QueryFilter{
			AllData: true,
		},
		Status:    sql.NullString{String: model.MeetingStatusNew, Valid: true},
		StartDate: sql.NullTime{Time: startDate, Valid: true},
	})
	if err != nil {
		return err
	}

	meets = helper.ExcludeDuplicateMeetings(meets)

	fmt.Printf("i found %d meetings by %v\n", len(meets), startDate)

	for _, meet := range meets {
		speaker, errMeet := s.studentRepository.GetList(ctx, &studentRepository.Query{
			QueryFilter: model.QueryFilter{
				AllData: true,
			},
			IDs: []int64{meet.SpeakerID},
		})
		if errMeet != nil {
			return errMeet
		}
		if len(speaker) == 0 {
			return fmt.Errorf("speaker with id %d not found", meet.SpeakerID)
		}

		listener, errMeet := s.studentRepository.GetList(ctx, &studentRepository.Query{
			QueryFilter: model.QueryFilter{
				AllData: true,
			},
			IDs: []int64{meet.ListenerID},
		})
		if errMeet != nil {
			return errMeet
		}
		if len(listener) == 0 {
			return fmt.Errorf("listener with id %d not found", meet.ListenerID)
		}

		unit, errMeet := s.unitRepository.GetList(ctx, &unitRepository.Query{
			QueryFilter: model.QueryFilter{
				AllData: true,
			},
			UnitIDs: []int64{meet.UnitID},
		})
		if errMeet != nil {
			return errMeet
		}
		if len(unit) == 0 {
			return fmt.Errorf("unit with id %d not found", meet.UnitID)
		}

		topic, errMeet := s.topicRepository.GetList(ctx, &topicRepository.Query{
			QueryFilter: model.QueryFilter{
				AllData: true,
			},
			UnitIDs:  []int64{meet.UnitID},
			TopicIDs: []int64{meet.TopicID},
		})
		if errMeet != nil {
			return errMeet
		}
		if len(topic) == 0 {
			return fmt.Errorf("topic with id %d not found", meet.TopicID)
		}

		err = s.sendNotificationToPartners(speaker[0], listener[0], unit[0], topic[0], meet.StartDate)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) sendNotificationToPartners(speaker *model.Student, listener *model.Student, unit *model.Unit, topic *model.Topic, startDate time.Time) error {
	speakerMsg, err := helper.GetNotification(listener, unit.Name, topic.Name, startDate, speaker.TelegramID, template.NotificationBeforeStart)
	if err != nil {
		fmt.Printf("error while getting notification message: %v\n", err)
		return err
	}

	listenerMsg, err := helper.GetNotification(speaker, unit.Name, topic.Name, startDate, listener.TelegramID, template.NotificationBeforeStart)
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
