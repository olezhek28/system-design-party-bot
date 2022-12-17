package processor

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/helper"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/model/command"
	"github.com/olezhek28/system-design-party-bot/internal/template"
)

func (s *Service) ListTopics(ctx context.Context, msg *model.TelegramMessage) (tgBotAPI.MessageConfig, error) {
	topics, err := s.topicRepository.GetTopicsByIDs(ctx, []int64{})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}
	if len(topics) == 0 {
		return tgBotAPI.NewMessage(msg.From.ID, "🚫 Что-то не нашёл в базе ни одной темы( Спроси у @olezhek28 в чём проблема."), nil
	}

	type topicInfo struct {
		ID          int64
		Name        string
		Description string
	}

	res := strings.Builder{}
	for _, topic := range topics {
		t, errTmpl := helper.ExecuteTemplate(template.TopicDescription, &topicInfo{
			ID:          topic.ID,
			Name:        topic.Name,
			Description: topic.Description,
		})
		if errTmpl != nil {
			return tgBotAPI.MessageConfig{}, errTmpl
		}

		res.WriteString(t + "\n")
	}

	t, err := helper.ExecuteTemplate(template.TopicGuidelines, struct {
		FirstName string
	}{
		FirstName: msg.From.FirstName,
	})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	res.WriteString(t)

	reply := tgBotAPI.NewMessage(msg.From.ID, res.String())

	var users []*model.Student
	users, err = s.studentRepository.GetStudentByTelegramChatIDs(ctx, []int64{msg.From.ID})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}
	if len(users) == 0 {
		return tgBotAPI.NewMessage(msg.From.ID, "Кажется ты не зарегистрирован:( Для этого нажми /"+command.Start), nil
	}

	listenerID := users[0].ID
	var speakerID int64
	flag := false
	if len(msg.Arguments) > 0 {
		flag, err = strconv.ParseBool(msg.Arguments[0])
		if err != nil {
			return tgBotAPI.MessageConfig{}, err
		}
	}

	if !flag {
		if len(msg.Arguments) == 2 {
			speakerID, err = strconv.ParseInt(msg.Arguments[1], 10, 64)
			if err != nil {
				return tgBotAPI.MessageConfig{}, err
			}
		}
	}

	reply.ReplyMarkup = getCreateMeetingKeyboard(topics, speakerID, listenerID)

	return reply, nil
}

func getCreateMeetingKeyboard(topics []*model.Topic, speakerID int64, listenerID int64) tgBotAPI.InlineKeyboardMarkup {
	var buttonsInfo []*model.TelegramButtonInfo
	for _, topic := range topics {
		buttonsInfo = append(buttonsInfo, &model.TelegramButtonInfo{
			Text: fmt.Sprintf("%d", topic.ID),
			Data: fmt.Sprintf("/%s %d %d %d", command.CreateMeeting, topic.ID, speakerID, listenerID),
		})
	}

	return helper.BuildKeyboard(buttonsInfo, 3)
}
