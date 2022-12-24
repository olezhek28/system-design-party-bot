package model

// SortOrder ...
const (
	OrderASC  = "ASC"
	OrderDESC = "DESC"
)

// QueryFilter ...
type QueryFilter struct {
	PageSize   uint64
	PageNumber uint64
	// Получить все данные, без пагинации
	AllData bool
	// Без сортировки
	WithoutSort bool
	// По какому полю сортировать
	SortField string
	// Тип сортировки (ASC, DESC)
	SortOrder string
}
