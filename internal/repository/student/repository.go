package student_repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i Repository -o ./mocks/ -s "_minimock.go"

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/pkg/db"
	"github.com/pkg/errors"
)

const tableName = "student"

// Repository ...
type Repository interface {
	Create(ctx context.Context, student *model.Student) error
	GetList(ctx context.Context, filter *Query) ([]*model.Student, error)
	Update(ctx context.Context, telegramID int64, updateStudent *model.UpdateStudent) error
	IsExist(ctx context.Context, telegramChatID int64) (bool, error)
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

func (r *repository) Create(ctx context.Context, student *model.Student) error {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("first_name", "last_name", "telegram_id", "telegram_username").
		Values(student.FirstName, student.LastName, student.TelegramID, student.TelegramUsername)

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "student_repository.Create",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, v...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetList(ctx context.Context, filter *Query) ([]*model.Student, error) {
	builder := sq.Select("id, first_name, last_name, telegram_id, telegram_username, timezone, created_at").
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
		Name:     "student_repository.GetList",
		QueryRaw: query,
	}

	var res []*model.Student
	err = r.db.DB().SelectContext(ctx, &res, q, v...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *repository) Update(ctx context.Context, telegramID int64, updateStudent *model.UpdateStudent) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"telegram_id": telegramID})

	if updateStudent.FirstName.Valid {
		builder = builder.Set("first_name", updateStudent.FirstName.String)
	}
	if updateStudent.LastName.Valid {
		builder = builder.Set("last_name", updateStudent.LastName.String)
	}
	if updateStudent.TelegramUsername.Valid {
		builder = builder.Set("telegram_username", updateStudent.TelegramUsername.String)
	}
	if updateStudent.Timezone.Valid {
		builder = builder.Set("timezone", updateStudent.Timezone.Int64)
	}

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "student_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, v...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) IsExist(ctx context.Context, telegramChatID int64) (bool, error) {
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
		Name:     "student_repository.IsExist",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, v...).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return id > 0, nil
}
