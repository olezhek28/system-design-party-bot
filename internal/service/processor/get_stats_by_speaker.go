package processor

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/template"
)

func (s *Service) GetStatsBySpeaker(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	if len(msg.Arguments) == 0 {
		return tgBotAPI.MessageConfig{}, fmt.Errorf("no arguments")
	}

	speakerID, err := strconv.ParseInt(msg.Arguments[0], 10, 64)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	meets, err := s.meetingRepository.GetSuccessMeetingBySpeaker(ctx, speakerID)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	var topicIds []int64
	for _, m := range meets {
		topicIds = append(topicIds, m.TopicID)
	}

	topicsInfo, err := s.topicRepository.GetTopicsByIDs(ctx, topicIds)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	topicMap := make(map[int64]*model.Topic)
	for _, topic := range topicsInfo {
		topicMap[topic.ID] = topic
	}

	stats := make(map[int64]int64, len(topicsInfo))
	for _, meet := range meets {
		stats[meet.TopicID]++
	}

	res := strings.Builder{}
	t, err := helper.ExecuteTemplate(template.SpeakerStatsIntroduction, struct {
		FirstName string
		LastName  string
	}{
		FirstName: msg.From.FirstName,
		LastName:  msg.From.LastName,
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	res.WriteString(t + "\n")

	for topicID, count := range stats {
		topic, ok := topicMap[topicID]
		if !ok {
			fmt.Errorf("topic with id %d not found\n", topicID)
			continue
		}

		t, errTmpl := helper.ExecuteTemplate(template.SpeakerStats, struct {
			TopicName string
			Count     int64
		}{
			TopicName: topic.Name,
			Count:     count,
		})
		if errTmpl != nil {
			return tgBotAPI.MessageConfig{}, errTmpl
		}

		res.WriteString(t + "\n")
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, res.String())
	return reply, nil
}