package processor

import (
	"context"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	studentRepository "github.com/olezhek28/system-design-party-bot/internal/repository/student"
	topicRepository "github.com/olezhek28/system-design-party-bot/internal/repository/topic"
)

func (s *Service) GetTopicStats(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	//meets, err := s.meetingRepository.GetList(ctx, &meetingRepository.Query{
	//	QueryFilter: model.QueryFilter{
	//		AllData: true,
	//	},
	//	Status: sql.NullString{String: model.MeetingStatusFinished, Valid: true},
	//})
	//if err != nil {
	//	return tgBotAPI.MessageConfig{}, err
	//}
	//
	//// unitID[topicID][speakerID] = count
	//// TODO сделать обертку для таких мап с интерфейсом добавления и со всякими проверками внутри
	//stats := make(map[int64]map[int64]map[int64]int64)
	//for _, meet := range meets {
	//	if _, ok := stats[meet.UnitID]; !ok {
	//		stats[meet.TopicID] = make(map[int64]map[int64]int64)
	//	}
	//	if _, ok := stats[meet.UnitID][meet.TopicID]; !ok {
	//		stats[meet.UnitID][meet.TopicID] = make(map[int64]int64)
	//	}
	//
	//	stats[meet.UnitID][meet.TopicID][meet.SpeakerID]++
	//}
	//
	//topicMap, err := s.getTopicMap(ctx, meets)
	//if err != nil {
	//	return tgBotAPI.MessageConfig{}, err
	//}
	//
	//speakerMap, err := s.getSpeakerMap(ctx, meets)
	//if err != nil {
	//	return tgBotAPI.MessageConfig{}, err
	//}
	//
	//res := strings.Builder{}
	//t, err := helper.ExecuteTemplate(template.TopicStatsIntroduction, nil)
	//if err != nil {
	//	return tgBotAPI.MessageConfig{}, err
	//}
	//
	//res.WriteString(t + "\n")
	//
	//for unitID, unitStats := range stats {
	//	for topicID, topicStats := range unitStats {
	//		topic := topicMap[unitID][topicID]
	//		t, err := helper.ExecuteTemplate(template.TopicStatsTopic, topic)
	//		if err != nil {
	//			return tgBotAPI.MessageConfig{}, err
	//		}
	//
	//		res.WriteString(t + "\n")
	//
	//		for speakerID, count := range topicStats {
	//
	//			speaker := speakerMap[speakerID]
	//			t, err := helper.ExecuteTemplate(template.TopicStatsSpeaker, map[string]interface{}{
	//				"Speaker": speaker,
	//				"Count":   count,
	//			})
	//			if err != nil {
	//				return tgBotAPI.MessageConfig{}, err
	//			}
	//
	//			res.WriteString(t + "\n")
	//		}
	//	}
	//}
	//
	//for tID, topicStats := range stats {
	//	topic, ok := topicMap[tID]
	//	if !ok {
	//		errors.Errorf("topic with id %d not found\n", tID)
	//		continue
	//	}
	//
	//	t, err = helper.ExecuteTemplate(template.TopicStatsHeader, struct {
	//		TopicName string
	//	}{
	//		TopicName: topic.Name,
	//	})
	//	if err != nil {
	//		return tgBotAPI.MessageConfig{}, err
	//	}
	//
	//	res.WriteString(t + "\n")
	//
	//	for speakerID, count := range topicStats {
	//		speaker, ok := speakerMap[speakerID]
	//		if !ok {
	//			errors.Errorf("speaker with id %d not found\n", speakerID)
	//			continue
	//		}
	//
	//		t, err = helper.ExecuteTemplate(template.TopicStats, struct {
	//			FirstName        string
	//			LastName         string
	//			TelegramUsername string
	//			Count            int64
	//		}{
	//			FirstName:        speaker.FirstName,
	//			LastName:         speaker.LastName,
	//			TelegramUsername: speaker.TelegramUsername,
	//			Count:            count,
	//		})
	//		if err != nil {
	//			return tgBotAPI.MessageConfig{}, err
	//		}
	//
	//		res.WriteString(t + "\n")
	//	}
	//}
	//
	//reply := tgBotAPI.NewMessage(msg.From.ID, res.String())
	//return reply, nil

	return tgBotAPI.MessageConfig{}, nil
}

func (s *Service) getTopicMap(ctx context.Context, meets []*model.Meeting) (map[int64]map[int64]*model.Topic, error) {
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
		return nil, err
	}

	topicMap := make(map[int64]map[int64]*model.Topic)
	for _, topic := range topicsInfo {
		if _, ok := topicMap[topic.UnitID]; !ok {
			topicMap[topic.UnitID] = make(map[int64]*model.Topic)
		}

		topicMap[topic.UnitID][topic.ID] = topic
	}

	return topicMap, nil
}

func (s *Service) getSpeakerMap(ctx context.Context, meets []*model.Meeting) (map[int64]*model.Student, error) {
	var speakerIds []int64
	for _, m := range meets {
		speakerIds = append(speakerIds, m.SpeakerID)
	}

	speakersInfo, err := s.studentRepository.GetList(ctx, &studentRepository.Query{
		QueryFilter: model.QueryFilter{
			AllData: true,
		},
		IDs: speakerIds,
	})
	if err != nil {
		return nil, err
	}

	speakerMap := make(map[int64]*model.Student)
	for _, speaker := range speakersInfo {
		speakerMap[speaker.ID] = speaker
	}

	return speakerMap, nil
}

func (s *Service) getListenerMap(ctx context.Context, meets []*model.Meeting) (map[int64]*model.Student, error) {
	var listenerIds []int64
	for _, m := range meets {
		listenerIds = append(listenerIds, m.ListenerID)
	}

	listenersInfo, err := s.studentRepository.GetList(ctx, &studentRepository.Query{
		QueryFilter: model.QueryFilter{
			AllData: true,
		},
		IDs: listenerIds,
	})
	if err != nil {
		return nil, err
	}

	listenerMap := make(map[int64]*model.Student)
	for _, speaker := range listenersInfo {
		listenerMap[speaker.ID] = speaker
	}

	return listenerMap, nil
}
