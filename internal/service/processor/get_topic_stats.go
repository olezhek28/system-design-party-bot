package processor

import (
	"context"
	"strings"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/template"
	"github.com/pkg/errors"
)

func (s *Service) GetTopicStats(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	meets, err := s.meetingRepository.GetSuccessMeeting(ctx)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	// topicID[speakerID] = count
	stats := make(map[int64]map[int64]int64)
	for _, meet := range meets {
		if _, ok := stats[meet.TopicID]; !ok {
			stats[meet.TopicID] = make(map[int64]int64)
		}

		stats[meet.TopicID][meet.SpeakerID]++
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
	t, err := helper.ExecuteTemplate(template.TopicStatsIntroduction, nil)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	res.WriteString(t + "\n")

	for tID, topicStats := range stats {
		topic, ok := topicMap[tID]
		if !ok {
			errors.Errorf("topic with id %d not found\n", tID)
			continue
		}

		t, err = helper.ExecuteTemplate(template.TopicStatsHeader, struct {
			TopicName string
		}{
			TopicName: topic.Name,
		})
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}

		res.WriteString(t + "\n")

		for speakerID, count := range topicStats {
			speaker, ok := speakerMap[speakerID]
			if !ok {
				errors.Errorf("speaker with id %d not found\n", speakerID)
				continue
			}

			t, err = helper.ExecuteTemplate(template.TopicStats, struct {
				FirstName string
				LastName  string
				Count     int64
			}{
				FirstName: speaker.FirstName,
				LastName:  speaker.LastName,
				Count:     count,
			})
			if err != nil {
				return tgBotAPI.MessageConfig{}, err
			}

			res.WriteString(t + "\n")
		}
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, res.String())
	return reply, nil
}

func (s *Service) getTopicMap(ctx context.Context, meets []*model.Meeting) (map[int64]*model.Topic, error) {
	var topicIds []int64
	for _, m := range meets {
		topicIds = append(topicIds, m.TopicID)
	}

	topicsInfo, err := s.topicRepository.GetTopicsByIDs(ctx, topicIds)
	if err != nil {
		return nil, err
	}

	topicMap := make(map[int64]*model.Topic)
	for _, topic := range topicsInfo {
		topicMap[topic.ID] = topic
	}

	return topicMap, nil
}

func (s *Service) getSpeakerMap(ctx context.Context, meets []*model.Meeting) (map[int64]*model.Student, error) {
	var speakerIds []int64
	for _, m := range meets {
		speakerIds = append(speakerIds, m.SpeakerID)
	}

	speakersInfo, err := s.studentRepository.GetStudentByIDs(ctx, speakerIds)
	if err != nil {
		return nil, err
	}

	speakerMap := make(map[int64]*model.Student)
	for _, speaker := range speakersInfo {
		speakerMap[speaker.ID] = speaker
	}

	return speakerMap, nil
}
