package processor

import (
	"context"
	"fmt"

	"github.com/olezhek28/system-design-party-bot/internal/model"
)

func (s *Service) FindSpeaker(ctx context.Context, msg *model.TelegramMessage) (string, error) {
	stats, err := s.meetingRepository.GetSpeakers(ctx)
	if err != nil {
		return "", err
	}
	if len(stats) == 0 {
		return "No speakers", nil
	}

	speaker := stats[0]
	minCount := stats[0].Count
	for _, stat := range stats {
		if stat.Count < minCount {
			minCount = stat.Count
			speaker = stat
		}
	}

	//res := strings.Builder{}
	//for _, stat := range stats {
	//	res.WriteString(fmt.Sprintf("%s %s (%s) TopicID: %d, Count: %d\n",
	//		stat.SpeakerFirstName, stat.SpeakerLastName, stat.SpeakerTelegramNickname, stat.TopicID, stat.Count))
	//}

	return fmt.Sprintf("%s %s (%s) TopicID: %d, Count: %d\n",
		speaker.SpeakerFirstName, speaker.SpeakerLastName, speaker.SpeakerTelegramNickname, speaker.TopicID, speaker.Count), nil
}
