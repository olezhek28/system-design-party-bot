package processor

import (
	"context"
	"fmt"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/model/command"
	studentRepository "github.com/olezhek28/system-design-party-bot/internal/repository/student"
)

func (s *Service) GetTimezone(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	student, err := s.studentRepository.GetList(ctx, &studentRepository.Query{
		TelegramIDs: []int64{msg.From.ID},
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}
	if len(student) == 0 {
		return tgBotAPI.NewMessage(msg.From.ID, "Кажется ты не зарегистрирован:( Для этого нажми /"+command.Start), nil
	}

	hours := student[0].Timezone.Int64 / 60
	minutes := student[0].Timezone.Int64 % 60

	var timezoneStr string
	if student[0].Timezone.Int64 > 0 {
		timezoneStr = fmt.Sprintf("+%d", hours)

		if minutes != 0 {
			timezoneStr = fmt.Sprintf("%s:%d", timezoneStr, minutes)
		}
	}
	if student[0].Timezone.Int64 < 0 {
		timezoneStr = fmt.Sprintf("-%d", hours)

		if minutes != 0 {
			timezoneStr = fmt.Sprintf("%s:%d", timezoneStr, minutes)
		}
	}

	reply := fmt.Sprintf("Твоя временная зона: UTC%s\n", timezoneStr)

	return tgBotAPI.NewMessage(msg.From.ID, reply), nil
}
