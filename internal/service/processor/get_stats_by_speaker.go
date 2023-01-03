package processor

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	meetingRepository "github.com/olezhek28/system-design-party-bot/internal/repository/meeting"
	studentRepository "github.com/olezhek28/system-design-party-bot/internal/repository/student"
	topicRepository "github.com/olezhek28/system-design-party-bot/internal/repository/topic"
	unitRepository "github.com/olezhek28/system-design-party-bot/internal/repository/unit"
	"github.com/olezhek28/system-design-party-bot/internal/template"
	"github.com/pkg/errors"
)

func (s *Service) GetStatsBySpeaker(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	if len(msg.Arguments) == 0 {
		return tgBotAPI.MessageConfig{}, errors.New("no arguments")
	}

	speakerID, err := strconv.ParseInt(msg.Arguments[0], 10, 64)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	meets, err := s.meetingRepository.GetList(ctx, &meetingRepository.Query{
		QueryFilter: model.QueryFilter{
			AllData: true,
		},
		Status:     sql.NullString{String: model.MeetingStatusFinished, Valid: true},
		SpeakerIDs: []int64{speakerID},
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	pairs := make(map[topicRepository.Pair]struct{}, len(meets))
	for _, meet := range meets {
		pairs[topicRepository.Pair{
			UnitID:  meet.UnitID,
			TopicID: meet.TopicID,
		}] = struct{}{}
	}

	topicsInfo, err := s.topicRepository.GetList(ctx, &topicRepository.Query{
		QueryFilter: model.QueryFilter{
			AllData: true,
		},
		Pairs: pairs,
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	topicMap := make(map[int64]map[int64]*model.Topic, len(topicsInfo))
	for _, topic := range topicsInfo {
		if _, ok := topicMap[topic.UnitID]; !ok {
			topicMap[topic.UnitID] = make(map[int64]*model.Topic)
		}

		topicMap[topic.UnitID][topic.ID] = topic
	}

	stats := make(map[int64]map[int64]int64, len(topicsInfo))
	for _, meet := range meets {
		if _, ok := stats[meet.UnitID]; !ok {
			stats[meet.UnitID] = make(map[int64]int64)
		}

		stats[meet.UnitID][meet.TopicID]++
	}

	unitIDs := make([]int64, 0, len(stats))
	for unitID := range stats {
		unitIDs = append(unitIDs, unitID)
	}

	unitInfo, err := s.unitRepository.GetList(ctx, &unitRepository.Query{
		QueryFilter: model.QueryFilter{
			AllData: true,
		},
		UnitIDs: unitIDs,
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	unitMap := make(map[int64]*model.Unit, len(unitInfo))
	for _, unit := range unitInfo {
		unitMap[unit.ID] = unit
	}

	speakers, err := s.studentRepository.GetList(ctx, &studentRepository.Query{
		QueryFilter: model.QueryFilter{
			AllData: true,
		},
		IDs: []int64{speakerID},
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}
	if len(speakers) == 0 {
		return tgBotAPI.NewMessage(msg.From.ID, "Что-то я не нашёл такого студента в базе:("), nil
	}

	res := strings.Builder{}
	t, err := helper.ExecuteTemplate(template.SpeakerStatsIntroduction, struct {
		FirstName string
		LastName  string
	}{
		FirstName: speakers[0].FirstName,
		LastName:  speakers[0].LastName,
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	res.WriteString(t + "\n")

	for unitID, topics := range stats {
		_, ok := topicMap[unitID]
		if !ok {
			return tgBotAPI.MessageConfig{}, errors.Errorf("unit with id %d not found\n", unitID)
		}

		t, err = helper.ExecuteTemplate(template.SpeakerStatsUnitIntroduction, struct {
			UnitName string
		}{
			UnitName: unitMap[unitID].Name,
		})
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}

		res.WriteString(t + "\n")

		for topicID, count := range topics {
			_, ok = topicMap[unitID][topicID]
			if !ok {
				return tgBotAPI.MessageConfig{}, errors.Errorf("topic with id %d not found\n", topicID)
			}

			t, err = helper.ExecuteTemplate(template.SpeakerStats, struct {
				TopicName string
				Count     int64
			}{
				TopicName: topicMap[unitID][topicID].Name,
				Count:     count,
			})
			if err != nil {
				return tgBotAPI.MessageConfig{}, err
			}

			res.WriteString(t + "\n")
		}
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, res.String())
	reply.ParseMode = tgBotAPI.ModeHTML
	return reply, nil
}
