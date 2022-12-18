package topic_repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i Repository -o ./mocks/ -s "_minimock.go"

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/pkg/db"
)

const tableName = "topic"

// Repository ...
type Repository interface {
	GetTopicsByIDs(ctx context.Context, ids []int64) ([]*model.Topic, error)
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

// GetTopicsByIDs ...
func (r *repository) GetTopicsByIDs(ctx context.Context, ids []int64) ([]*model.Topic, error) {
	builder := sq.Select("id, name, description, link, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		OrderBy("id ASC")

	if len(ids) > 0 {
		builder = builder.Where(sq.Eq{"id": ids})
	}

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "topic_repository.GetTopicList",
		QueryRaw: query,
	}

	var res []*model.Topic
	err = r.db.DB().SelectContext(ctx, &res, q, v...)
	if err != nil {
		return nil, err
	}

	return res, nil
}
