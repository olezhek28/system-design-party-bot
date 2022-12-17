package helper

import (
	"strings"

	"github.com/olezhek28/system-design-party-bot/internal/model"
)

func SplitSlice(data []*model.TelegramButtonInfo, chunkSize int) [][]*model.TelegramButtonInfo {
	var chunks [][]*model.TelegramButtonInfo

	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize

		if end > len(data) {
			end = len(data)
		}

		chunks = append(chunks, data[i:end])
	}

	return chunks
}

func GetMonthList() map[int64]string {
	return map[int64]string{
		1:  "Январь",
		2:  "Февраль",
		3:  "Март",
		4:  "Апрель",
		5:  "Май",
		6:  "Июнь",
		7:  "Июль",
		8:  "Август",
		9:  "Сентябрь",
		10: "Октябрь",
		11: "Ноябрь",
		12: "Декабрь",
	}
}

func GetDaysInMonth(year int64, month int64) []int64 {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}
	case 4, 6, 9, 11:
		return []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	case 2:
		if isLeapYear(year) {
			return []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}
		}
		return []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28}
	default:
		return []int64{}
	}
}

func GetHours() []int64 {
	hours := make([]int64, 0, 24)
	for i := 0; i < 24; i++ {
		hours = append(hours, int64(i))
	}

	return hours
}

func GetMinutes() []int64 {
	minutes := make([]int64, 0, 6)
	for i := 0; i < 60; i += 10 {
		minutes = append(minutes, int64(i))
	}

	return minutes
}

func isLeapYear(year int64) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func SliceToString(data []string) string {
	str := strings.Builder{}
	for _, arg := range data {
		str.WriteString(arg)
		str.WriteString(" ")
	}

	return str.String()
}

func GetInexperiencedSpeaker(stats []*model.Stats) int64 {
	if len(stats) == 0 {
		return 0
	}

	id := stats[0].SpeakerID
	minCount := stats[0].Count
	for _, stat := range stats {
		if stat.Count < minCount {
			minCount = stat.Count
			id = stat.SpeakerID
		}
	}

	return id
}
