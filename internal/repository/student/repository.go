package student_repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i Repository -o ./mocks/ -s "_minimock.go"

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/pkg/db"
)

const tableName = "student"

// Repository ...
type Repository interface {
	GetStudentByIDs(ctx context.Context, ids []int64) ([]*model.Student, error)
}

type repository struct {
	db db.Client
}

// NewRepository ...
func NewRepository(db db.Client) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetStudentByIDs(ctx context.Context, ids []int64) ([]*model.Student, error) {
	builder := sq.Select("id, first_name, last_name, telegram_chat_id, telegram_username, created_at").
		PlaceholderFormat(sq.Dollar).
		From(tableName)

	if len(ids) > 0 {
		builder = builder.Where(sq.Eq{"id": ids})
	}

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "topic_repository.GetStudentByIDs",
		QueryRaw: query,
	}

	var res []*model.Student
	err = r.db.DB().SelectContext(ctx, &res, q, v...)
	if err != nil {
		return nil, err
	}

	return res, nil
}
