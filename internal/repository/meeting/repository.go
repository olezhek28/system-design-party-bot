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
)

// Repository ...
type Repository interface {
	GetSpeakers(ctx context.Context) ([]*model.Stats, error)
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

// select topic_id, s.first_name, s.last_name, count(*) from meeting m join student s on m.speaker_id = s.id group by topic_id, s.first_name, s.last_name,status having status='success';
// GetSpeakers ...
func (r *repository) GetSpeakers(ctx context.Context) ([]*model.Stats, error) {
	builder := sq.Select("m.topic_id, s.first_name, s.last_name, s.telegram_username, count(*) ").
		PlaceholderFormat(sq.Dollar).
		From(meetingTable + " m").
		Join(studentTable + " s on m.speaker_id=s.id").
		GroupBy("m.topic_id, s.first_name, s.last_name, s.telegram_username, m.status").
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
