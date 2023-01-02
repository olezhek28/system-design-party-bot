package scheduler

import (
	"github.com/olezhek28/system-design-party-bot/internal/pkg/http/telegram"
	meetingRepository "github.com/olezhek28/system-design-party-bot/internal/repository/meeting"
	studentRepository "github.com/olezhek28/system-design-party-bot/internal/repository/student"
	topicRepository "github.com/olezhek28/system-design-party-bot/internal/repository/topic"
	unitRepository "github.com/olezhek28/system-design-party-bot/internal/repository/unit"
)

// Service ...
type Service struct {
	telegramClient telegram.Client

	meetingRepository meetingRepository.Repository
	topicRepository   topicRepository.Repository
	unitRepository    unitRepository.Repository
	studentRepository studentRepository.Repository
}

// NewService ...
func NewService(
	telegramClient telegram.Client,
	meetingRepository meetingRepository.Repository,
	topicRepository topicRepository.Repository,
	unitRepository unitRepository.Repository,
	studentRepository studentRepository.Repository,
) *Service {
	return &Service{
		telegramClient:    telegramClient,
		meetingRepository: meetingRepository,
		topicRepository:   topicRepository,
		unitRepository:    unitRepository,
		studentRepository: studentRepository,
	}
}
