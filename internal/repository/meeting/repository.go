package meeting_repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i Repository -o ./mocks/ -s "_minimock.go"

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/pkg/db"
)

const (
	meetingTable = "meeting"
	studentTable = "student"
	topicTable   = "topic"
)

// Repository ...
type Repository interface {
	GetSpeakers(ctx context.Context, topicID int64) ([]*model.Stats, error)
	GetSuccessMeetingBySpeaker(ctx context.Context, speakerID int64) ([]*model.Meeting, error)
	GetSuccessMeeting(ctx context.Context) ([]*model.Meeting, error)
}

type repository struct {
	db db.Client
}

// NewRepository ...
func NewRepository(db db.Client) *repository {
	return &repository{
		db: db,
	}
}

// GetSpeakers ...
func (r *repository) GetSpeakers(ctx context.Context, topicID int64) ([]*model.Stats, error) {
	builder := sq.Select("t.name, s.id, s.first_name, s.last_name, s.telegram_username, count(*) ").
		PlaceholderFormat(sq.Dollar).
		From(meetingTable + " m").
		Join(studentTable + " s on m.speaker_id=s.id").
		Join(topicTable + " t on m.topic_id=t.id").
		Where(sq.Eq{"m.topic_id": topicID}).
		GroupBy("t.name, s.id, s.first_name, s.last_name, s.telegram_username, m.status").
		Having("m.status='success'")

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "meeting_repository.GetSpeakers",
		QueryRaw: query,
	}

	var res []*model.Stats
	err = r.db.DB().SelectContext(ctx, &res, q, v...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *repository) GetSuccessMeetingBySpeaker(ctx context.Context, speakerID int64) ([]*model.Meeting, error) {
	builder := sq.Select("id, topic_id, status, start_date, speaker_id, listener_id, created_at").
		PlaceholderFormat(sq.Dollar).
		From(meetingTable).
		Where(sq.Eq{"speaker_id": speakerID}).
		Where(sq.Eq{"status": "success"})

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "meeting_repository.GetSuccessMeetingBySpeaker",
		QueryRaw: query,
	}

	var res []*model.Meeting
	err = r.db.DB().SelectContext(ctx, &res, q, v...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *repository) GetSuccessMeeting(ctx context.Context) ([]*model.Meeting, error) {
	builder := sq.Select("id, topic_id, status, start_date, speaker_id, listener_id, created_at").
		PlaceholderFormat(sq.Dollar).
		From(meetingTable).
		Where(sq.Eq{"status": "success"})

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "meeting_repository.GetSuccessMeetingByTopic",
		QueryRaw: query,
	}

	var res []*model.Meeting
	err = r.db.DB().SelectContext(ctx, &res, q, v...)
	if err != nil {
		return nil, err
	}

	return res, nil
}
