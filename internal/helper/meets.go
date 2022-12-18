package helper

import "github.com/olezhek28/system-design-party-bot/internal/model"

// TODO refactor, del n*n complexity
func ExcludeDuplicateMeetings(meetings []*model.Meeting) []*model.Meeting {
	res := make([]*model.Meeting, 0, len(meetings))
	for i := 0; i < len(meetings); i++ {
		isDuplicate := false
		for j := i + 1; j < len(meetings); j++ {
			if meetings[i].SpeakerID == meetings[j].ListenerID &&
				meetings[i].ListenerID == meetings[j].SpeakerID &&
				meetings[i].TopicID == meetings[j].TopicID &&
				meetings[i].StartDate == meetings[j].StartDate {
				isDuplicate = true
				break
			}
		}

		if !isDuplicate {
			res = append(res, meetings[i])
		}
	}

	return res
}
