package processor

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"unicode/utf8"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/model/command"
	"github.com/olezhek28/system-design-party-bot/internal/template"
)

func (s *Service) GetStudents(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	students, err := s.studentRepository.GetStudentList(ctx)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	t, err := helper.ExecuteTemplate(template.StudentDescription, struct {
		Emoji string
	}{
		Emoji: model.GetEmoji(model.PlantsEmojis),
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	reply := tgBotAPI.NewMessage(msg.From.ID, t)
	reply.ReplyMarkup = getStudentsKeyboard(students)

	return reply, nil
}

func getStudentsKeyboard(students []*model.Student) tgBotAPI.InlineKeyboardMarkup {
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

	return helper.BuildKeyboard(buttonsInfo)
}

func getStudentText(student *model.Student) (string, error) {
	rand.Seed(time.Now().UnixNano())
	emoji := []rune(model.VegetablesEmojis)[rand.Intn(utf8.RuneCountInString(model.VegetablesEmojis))]

	t, err := helper.ExecuteTemplate(template.StudentInfo, struct {
		Emoji     string
		FirstName string
		LastName  string
	}{
		Emoji:     string(emoji),
		FirstName: student.FirstName,
		LastName:  student.LastName,
	})
	if err != nil {
		return "", err
	}

	return t, nil
}
