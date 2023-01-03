package processor

import (
	"context"
	"database/sql"
	"strings"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/model/command"
	meetingRepository "github.com/olezhek28/system-design-party-bot/internal/repository/meeting"
	studentRepository "github.com/olezhek28/system-design-party-bot/internal/repository/student"
	"github.com/olezhek28/system-design-party-bot/internal/template"
)

func (s *Service) GetSocialConnections(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	user, err := s.studentRepository.GetList(ctx, &studentRepository.Query{
		QueryFilter: model.QueryFilter{
			AllData: true,
		},
		TelegramIDs: []int64{msg.From.ID},
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}
	if len(user) == 0 {
		return tgBotAPI.NewMessage(msg.From.ID, "Кажется ты не зарегистрирован:( Для этого нажми /"+command.Start), nil
	}

	meets, err := s.meetingRepository.GetList(ctx, &meetingRepository.Query{
		QueryFilter: model.QueryFilter{
			AllData: true,
		},
		Status:     sql.NullString{String: model.MeetingStatusFinished, Valid: true},
		SpeakerIDs: []int64{user[0].ID},
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	students, err := s.studentRepository.GetList(ctx, &studentRepository.Query{
		QueryFilter: model.QueryFilter{
			AllData: true,
		},
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	studentMap := make(map[int64]*model.Student, len(students))
	for _, student := range students {
		studentMap[student.ID] = student
	}

	partners := make(map[int64]*model.Student)
	for _, meet := range meets {
		partners[meet.ListenerID] = studentMap[meet.ListenerID]
	}

	res := strings.Builder{}
	t, err := helper.ExecuteTemplate(template.SocialOwnConnectionDescription, struct {
		Emoji string
	}{
		Emoji: model.GetEmoji(model.FoodEmojis),
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	res.WriteString(t + "\n")

	for _, partner := range partners {
		t, err = helper.ExecuteTemplate(template.SocialConnection, struct {
			PartnerFirstName        string
			PartnerLastName         string
			PartnerTelegramUsername string
		}{
			PartnerFirstName:        partner.FirstName,
			PartnerLastName:         partner.LastName,
			PartnerTelegramUsername: partner.TelegramUsername,
		})
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}

		res.WriteString(t)
	}

	for _, student := range students {
		if _, ok := partners[student.ID]; ok {
			continue
		}

		t, err = helper.ExecuteTemplate(template.SocialNotConnection, struct {
			PartnerFirstName        string
			PartnerLastName         string
			PartnerTelegramUsername string
		}{
			PartnerFirstName:        student.FirstName,
			PartnerLastName:         student.LastName,
			PartnerTelegramUsername: student.TelegramUsername,
		})
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}

		res.WriteString(t)
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, res.String())

	return reply, nil
}
