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
	GetStudentByTelegramChatIDs(ctx context.Context, ids []int64) ([]*model.Student, error)
	CreateStudent(ctx context.Context, student *model.Student) error
	IsExistStudent(ctx context.Context, telegramChatID int64) (bool, error)
	GetRandomStudent(ctx context.Context, excludeSpeakerTelegramID int64) (*model.Student, error)
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
	builder := sq.Select("id, first_name, last_name, telegram_id, telegram_username, created_at").
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

func (r *repository) GetStudentByTelegramChatIDs(ctx context.Context, ids []int64) ([]*model.Student, error) {
	builder := sq.Select("id, first_name, last_name, telegram_id, telegram_username, created_at").
		PlaceholderFormat(sq.Dollar).
		From(tableName)

	if len(ids) > 0 {
		builder = builder.Where(sq.Eq{"telegram_id": ids})
	}

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "topic_repository.GetStudentByTelegramChatIDs",
		QueryRaw: query,
	}

	var res []*model.Student
	err = r.db.DB().SelectContext(ctx, &res, q, v...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *repository) CreateStudent(ctx context.Context, student *model.Student) error {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("first_name", "last_name", "telegram_id", "telegram_username").
		Values(student.FirstName, student.LastName, student.TelegramID, student.TelegramUsername)

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "topic_repository.CreateStudent",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, v...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) IsExistStudent(ctx context.Context, telegramChatID int64) (bool, error) {
	builder := sq.Select("id").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{"telegram_id": telegramChatID}).
		Limit(1)

	query, v, err := builder.ToSql()
	if err != nil {
		return false, err
	}

	q := db.Query{
		Name:     "topic_repository.IsExistStudent",
		QueryRaw: query,
	}

	var ids []int64
	err = r.db.DB().SelectContext(ctx, &ids, q, v...)
	if err != nil {
		return false, err
	}

	return len(ids) > 0, nil
}

func (r *repository) GetRandomStudent(ctx context.Context, excludeSpeakerTelegramID int64) (*model.Student, error) {
	builder := sq.Select("id, first_name, last_name, telegram_id, telegram_username, created_at").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		OrderBy("RANDOM()").
		Where(sq.NotEq{"telegram_id": excludeSpeakerTelegramID}).
		Limit(1)

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "topic_repository.GetRandomStudent",
		QueryRaw: query,
	}

	var res []*model.Student
	err = r.db.DB().SelectContext(ctx, &res, q, v...)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}

	return res[0], nil
}
