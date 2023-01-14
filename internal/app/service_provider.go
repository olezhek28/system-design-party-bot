package app

import (
	"context"
	"log"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/system-design-party-bot/internal/config"
	"github.com/olezhek28/system-design-party-bot/internal/pkg/db"
	"github.com/olezhek28/system-design-party-bot/internal/pkg/http/telegram"
	meetingRepository "github.com/olezhek28/system-design-party-bot/internal/repository/meeting"
	studentRepository "github.com/olezhek28/system-design-party-bot/internal/repository/student"
	topicRepository "github.com/olezhek28/system-design-party-bot/internal/repository/topic"
	unitRepository "github.com/olezhek28/system-design-party-bot/internal/repository/unit"
	"github.com/olezhek28/system-design-party-bot/internal/service/processor"
	"github.com/olezhek28/system-design-party-bot/internal/service/scheduler"
)

type serviceProvider struct {
	db db.Client

	telegramClient telegram.Client

	meetingRepository meetingRepository.Repository
	topicRepository   topicRepository.Repository
	unitRepository    unitRepository.Repository
	studentRepository studentRepository.Repository

	processorService *processor.Service
	schedulerService *scheduler.Service
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GetDB(ctx context.Context) db.Client {
	if s.db == nil {
		cfg, err := config.GetDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		db, err := db.NewClient(ctx, cfg.DSN())
		if err != nil {
			log.Fatalf("failed to creating new db client: %s", err.Error())
		}

		s.db = db
	}

	return s.db
}

func (s *serviceProvider) GetTelegramClient() telegram.Client {
	if s.telegramClient == nil {
		cfg, err := config.GetTelegramBotConfig()
		if err != nil {
			log.Fatalf("failed to get telegram config: %s", err.Error())
		}

		bot, err := tgBotAPI.NewBotAPI(cfg.TelegramToken())
		if err != nil {
			log.Fatalf("failed to creating new tg client: %s", err.Error())
		}

		s.telegramClient = telegram.NewClient(bot, cfg)
	}

	return s.telegramClient
}

func (s *serviceProvider) GetMeetingRepository(ctx context.Context) meetingRepository.Repository {
	if s.meetingRepository == nil {
		s.meetingRepository = meetingRepository.NewRepository(s.GetDB(ctx))
	}

	return s.meetingRepository
}

func (s *serviceProvider) GetTopicRepository(ctx context.Context) topicRepository.Repository {
	if s.topicRepository == nil {
		s.topicRepository = topicRepository.NewRepository(s.GetDB(ctx))
	}

	return s.topicRepository
}

func (s *serviceProvider) GetUnitRepository(ctx context.Context) unitRepository.Repository {
	if s.unitRepository == nil {
		s.unitRepository = unitRepository.NewRepository(s.GetDB(ctx))
	}

	return s.unitRepository
}

func (s *serviceProvider) GetStudentRepository(ctx context.Context) studentRepository.Repository {
	if s.studentRepository == nil {
		s.studentRepository = studentRepository.NewRepository(s.GetDB(ctx))
	}

	return s.studentRepository
}

func (s *serviceProvider) GetProcessorService(ctx context.Context) *processor.Service {
	if s.processorService == nil {
		s.processorService = processor.NewService(
			s.GetTelegramClient(),
			s.GetMeetingRepository(ctx),
			s.GetTopicRepository(ctx),
			s.GetUnitRepository(ctx),
			s.GetStudentRepository(ctx),
		)
	}

	return s.processorService
}

func (s *serviceProvider) GetSchedulerService(ctx context.Context) *scheduler.Service {
	if s.schedulerService == nil {
		s.schedulerService = scheduler.NewService(
			s.GetTelegramClient(),
			s.GetMeetingRepository(ctx),
			s.GetTopicRepository(ctx),
			s.GetUnitRepository(ctx),
			s.GetStudentRepository(ctx),
		)
	}

	return s.schedulerService
}
