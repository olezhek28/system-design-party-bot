package processor

import (
	"context"
	"fmt"
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
	reply.ReplyMarkup = getTopicsKeyboard(topics)

	return reply, nil
}

func getTopicsKeyboard(topics []*model.Topic) tgBotAPI.InlineKeyboardMarkup {
	var buttonsInfo []*model.TelegramButtonInfo
	for _, topic := range topics {
		buttonsInfo = append(buttonsInfo, &model.TelegramButtonInfo{
			Text: fmt.Sprintf("%d", topic.ID),
			Data: fmt.Sprintf("/%s %d", command.FindSpeaker, topic.ID),
		})
	}

	return helper.BuildKeyboard(buttonsInfo)
}
