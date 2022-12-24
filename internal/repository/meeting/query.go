package meeting_repository

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/olezhek28/system-design-party-bot/internal/model"
)

const (
	defaultPageSize = 50
)

// Query ...
type Query struct {
	model.QueryFilter

	IDs               []int64
	TopicIDs          []int64
	Status            sql.NullString
	StartDate         sql.NullTime
	BeginStartDate    sql.NullTime
	EndStartDate      sql.NullTime
	SpeakerIDs        []int64
	ExcludeSpeakerIDs []int64
	ListenerIDs       []int64
	CreatedAt         sql.NullTime
	BeginCreatedAt    sql.NullTime
	EndCreatedAt      sql.NullTime
}

func (q *Query) executeFilter(builder sq.SelectBuilder) sq.SelectBuilder {
	if len(q.IDs) > 0 {
		builder = builder.Where(sq.Eq{"id": q.IDs})
	}
	if len(q.TopicIDs) > 0 {
		builder = builder.Where(sq.Eq{"topic_id": q.TopicIDs})
	}
	if q.Status.Valid {
		builder = builder.Where(sq.Eq{"status": q.Status.String})
	}
	if q.StartDate.Valid {
		builder = builder.Where(sq.Eq{"start_date": q.StartDate.Time})
	}
	if q.BeginStartDate.Valid {
		builder = builder.Where(sq.GtOrEq{"start_date": q.BeginStartDate.Time})
	}
	if q.EndStartDate.Valid {
		builder = builder.Where(sq.Lt{"start_date": q.EndStartDate.Time})
	}
	if len(q.SpeakerIDs) > 0 {
		builder = builder.Where(sq.Eq{"speaker_id": q.SpeakerIDs})
	}
	if len(q.ExcludeSpeakerIDs) > 0 {
		builder = builder.Where(sq.NotEq{"speaker_id": q.ExcludeSpeakerIDs})
	}
	if len(q.ListenerIDs) > 0 {
		builder = builder.Where(sq.Eq{"listener_id": q.ListenerIDs})
	}
	if q.CreatedAt.Valid {
		builder = builder.Where(sq.Eq{"created_at": q.CreatedAt.Time})
	}
	if q.BeginCreatedAt.Valid {
		builder = builder.Where(sq.GtOrEq{"created_at": q.BeginCreatedAt.Time})
	}
	if q.EndCreatedAt.Valid {
		builder = builder.Where(sq.Lt{"created_at": q.EndCreatedAt.Time})
	}

	if !q.AllData {
		q.PageSize = q.getLimit()
		builder = builder.Limit(q.PageSize)

		builder = builder.Offset(q.getOffset())
	}

	if !q.WithoutSort {
		builder = builder.OrderBy(q.getSortMode())
	}

	return builder
}

func (q *Query) getSortMode() string {
	sortField := q.getSortField()
	sortType := q.getSortOrder()
	return fmt.Sprintf("%s %s", sortField, sortType)
}

func (q *Query) getLimit() uint64 {
	if q.PageSize == 0 {
		return defaultPageSize
	}

	return q.PageSize
}

func (q *Query) getOffset() uint64 {
	if q.PageNumber == 0 {
		return 0
	}

	return (q.PageNumber - 1) * q.PageSize
}

func (q *Query) getSortField() string {
	if q.SortField == "" {
		return "id"
	}

	return q.SortField
}

func (q *Query) getSortOrder() string {
	if q.SortOrder == model.OrderASC || q.SortOrder == model.OrderDESC {
		return q.SortOrder
	}

	return model.OrderASC
}
