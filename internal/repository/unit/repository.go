package unit_repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i Repository -o ./mocks/ -s "_minimock.go"

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/pkg/db"
)

const tableName = "unit"

// Repository ...
type Repository interface {
	GetList(ctx context.Context, filter *Query) ([]*model.Unit, error)
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

func (r *repository) GetList(ctx context.Context, filter *Query) ([]*model.Unit, error) {
	builder := sq.Select("id, name, description, link, created_at, updated_at").
		PlaceholderFormat(sq.Dollar).
		From(tableName)

	if filter != nil {
		builder = filter.executeFilter(builder)
	}

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "unit_repository.GetList",
		QueryRaw: query,
	}

	var res []*model.Unit
	err = r.db.DB().SelectContext(ctx, &res, q, v...)
	if err != nil {
		return nil, err
	}

	return res, nil
}
