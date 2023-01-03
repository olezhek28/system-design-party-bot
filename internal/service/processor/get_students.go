package processor

import (
	"context"
	"fmt"
	"strconv"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/model/command"
	studentRepository "github.com/olezhek28/system-design-party-bot/internal/repository/student"
	"github.com/olezhek28/system-design-party-bot/internal/template"
)

func (s *Service) GetStudents(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	students, err := s.studentRepository.GetList(ctx, &studentRepository.Query{
		QueryFilter: model.QueryFilter{
			AllData: true,
		},
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	autoChoiceSpeaker := true
	if len(msg.Arguments) == 1 {
		autoChoiceSpeaker, err = strconv.ParseBool(msg.Arguments[0])
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}
	}

	tmpl := template.StudentStatsDescription
	if !autoChoiceSpeaker {
		tmpl = template.StudentCreateMeetingDescription
	}

	t, err := helper.ExecuteTemplate(tmpl, struct {
		Emoji string
	}{
		Emoji: model.GetEmoji(model.PlantsEmojis),
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, t)
	reply.ReplyMarkup = getStudentsWithStatsKeyboard(students)
	if !autoChoiceSpeaker {
		reply.ReplyMarkup = getStudentsWithListUnitsKeyboard(excludeStudents(students, []int64{msg.From.ID}), msg.Arguments)
	}

	return reply, nil
}

func getStudentsWithStatsKeyboard(students []*model.Student) tgBotAPI.InlineKeyboardMarkup {
	var buttonsInfo []*model.TelegramButtonInfo
	for _, st := range students {
		text, err := getStudentText(st)
		if err != nil {
			fmt.Printf("error while getting student text: %v", err)
			continue
		}

		buttonsInfo = append(buttonsInfo, &model.TelegramButtonInfo{
			Text: text,
			Data: fmt.Sprintf("/%s %d", command.GetStatsBySpeaker, st.ID),
		})
	}

	return helper.BuildKeyboard(buttonsInfo, 2)
}

func getStudentsWithListUnitsKeyboard(students []*model.Student, args []string) tgBotAPI.InlineKeyboardMarkup {
	var buttonsInfo []*model.TelegramButtonInfo
	for _, st := range students {
		text, err := getStudentText(st)
		if err != nil {
			fmt.Printf("error while getting student text: %v", err)
			continue
		}

		buttonsInfo = append(buttonsInfo, &model.TelegramButtonInfo{
			Text: text,
			Data: fmt.Sprintf("/%s %s %d", command.ListUnits, helper.SliceToString(args), st.ID),
		})
	}

	return helper.BuildKeyboard(buttonsInfo, 2)
}

func getStudentText(student *model.Student) (string, error) {
	t, err := helper.ExecuteTemplate(template.StudentInfo, struct {
		Emoji     string
		FirstName string
		LastName  string
	}{
		Emoji:     model.GetEmoji(model.VegetablesEmojis),
		FirstName: student.FirstName,
		LastName:  student.LastName,
	})
	if err != nil {
		return "", err
	}

	return t, nil
}

func excludeStudents(students []*model.Student, studentTelegramIDs []int64) []*model.Student {
	excludeMap := make(map[int64]struct{}, len(studentTelegramIDs))
	for _, id := range studentTelegramIDs {
		excludeMap[id] = struct{}{}
	}

	res := make([]*model.Student, 0, len(students))
	for _, st := range students {
		if _, ok := excludeMap[st.TelegramID]; !ok {
			res = append(res, st)
		}
	}

	return res
}
