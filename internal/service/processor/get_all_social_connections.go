package processor

import (
	"context"
	"database/sql"
	"strings"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	meetingRepository "github.com/olezhek28/system-design-party-bot/internal/repository/meeting"
	studentRepository "github.com/olezhek28/system-design-party-bot/internal/repository/student"
	"github.com/olezhek28/system-design-party-bot/internal/template"
	"github.com/pkg/errors"
)

func (s *Service) GetAllSocialConnections(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	meets, err := s.meetingRepository.GetList(ctx, &meetingRepository.Query{
		QueryFilter: model.QueryFilter{
			AllData: true,
		},
		Status: sql.NullString{String: model.MeetingStatusFinished, Valid: true},
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

	stats := make(map[int64]map[int64]*model.Student)
	for _, meet := range meets {
		if _, ok := stats[meet.SpeakerID]; !ok {
			stats[meet.SpeakerID] = make(map[int64]*model.Student)
		}

		stats[meet.SpeakerID][meet.ListenerID] = studentMap[meet.ListenerID]
	}

	res := strings.Builder{}
	t, err := helper.ExecuteTemplate(template.SocialConnectionDescription, struct {
		Emoji string
	}{
		Emoji: model.GetEmoji(model.FoodEmojis),
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	res.WriteString(t + "\n")

	for studentID, partners := range stats {
		studentInfo, ok := studentMap[studentID]
		if !ok {
			return tgBotAPI.MessageConfig{}, errors.New("student not found")
		}

		t, err = helper.ExecuteTemplate(template.SocialConnectionStudentName, struct {
			StudentFirstName        string
			StudentLastName         string
			StudentTelegramUsername string
			Emoji                   string
		}{
			StudentFirstName:        studentInfo.FirstName,
			StudentLastName:         studentInfo.LastName,
			StudentTelegramUsername: studentInfo.TelegramUsername,
			Emoji:                   model.GetEmoji(model.AnimalsEmojis),
		})
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}

		res.WriteString(t)

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
			if _, ok = partners[student.ID]; ok {
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

		res.WriteString("\n")
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, res.String())

	return reply, nil
}
