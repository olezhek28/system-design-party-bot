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
	unitID, err := strconv.ParseInt(msg.Arguments[0], 10, 64)
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}

	topics, err := s.topicRepository.GetTopicsByIDs(ctx, []int64{unitID}, []int64{})
	if err != nil {
		return tgBotAPI.MessageConfig{}, err
	}
	if len(topics) == 0 {
		return tgBotAPI.NewMessage(msg.From.ID, "🚫 Что-то не нашёл в базе ни одной темы( Спроси у @olezhek28 в чём проблема."), nil
	}

	res := strings.Builder{}
	for _, topic := range topics {
		t, errTmpl := helper.ExecuteTemplate(template.TopicDescription, struct {
			ID          int64
			Name        string
			Description string
		}{
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
	reply.ReplyMarkup = getCreateMeetingKeyboard(topics, msg.Arguments)

	return reply, nil
}

func getCreateMeetingKeyboard(topics []*model.Topic, args []string) tgBotAPI.InlineKeyboardMarkup {
	var buttonsInfo []*model.TelegramButtonInfo
	for _, topic := range topics {
		buttonsInfo = append(buttonsInfo, &model.TelegramButtonInfo{
			Text: fmt.Sprintf("%d", topic.ID),
			Data: fmt.Sprintf("/%s %d %s", command.CreateMeeting, topic.ID, helper.SliceToString(args)),
		})
	}

	return helper.BuildKeyboard(buttonsInfo, 3)
}
