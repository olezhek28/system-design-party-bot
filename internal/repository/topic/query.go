package topic_repository

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/olezhek28/system-design-party-bot/internal/model"
)

const (
	defaultPageSize = 50
)

type Pair struct {
	UnitID  int64
	TopicID int64
}

// Query ...
type Query struct {
	model.QueryFilter

	UnitIDs  []int64
	TopicIDs []int64
	Pairs    map[Pair]struct{}
}

func (q *Query) executeFilter(builder sq.SelectBuilder) sq.SelectBuilder {
	if len(q.UnitIDs) > 0 {
		builder = builder.Where(sq.Eq{"unit_id": q.UnitIDs})
	}
	if len(q.TopicIDs) > 0 {
		builder = builder.Where(sq.Eq{"id": q.TopicIDs})
	}
	if len(q.Pairs) > 0 {
		var andConditions []sq.And
		for pair := range q.Pairs {
			andConditions = append(andConditions, sq.And{
				sq.Eq{"unit_id": pair.UnitID},
				sq.Eq{"id": pair.TopicID},
			})
			//builder = builder.Where(sq.And{sq.Eq{"unit_id": pair.UnitID}, sq.Eq{"id": pair.TopicID}})
		}

		orConditions := sq.Or{}
		for _, condition := range andConditions {
			orConditions = append(orConditions, condition)
		}

		builder = builder.Where(orConditions)
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
