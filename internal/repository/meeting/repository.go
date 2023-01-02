package meeting_repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i Repository -o ./mocks/ -s "_minimock.go"

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/pkg/db"
)

const tableName = "meeting"

// Repository ...
type Repository interface {
	Create(ctx context.Context, meeting *model.Meeting) (int64, error)
	Get(ctx context.Context, id int64) (*model.Meeting, error)
	GetList(ctx context.Context, filter *Query) ([]*model.Meeting, error)
	Update(ctx context.Context, ids []int64, updateMeeting *model.UpdateMeeting) error
	GetSpeakerCountByTopic(ctx context.Context, topicID int64, speakerID int64) (int64, error)
	GetSpeakersStats(ctx context.Context, unitID int64, topicID int64, excludeSpeakerID int64) ([]*model.Stats, error)
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

func (r *repository) Create(ctx context.Context, meeting *model.Meeting) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("unit_id", "topic_id", "status", "start_date", "speaker_id", "listener_id").
		Values(meeting.UnitID, meeting.TopicID, meeting.Status, meeting.StartDate, meeting.SpeakerID, meeting.ListenerID).
		Suffix("RETURNING id")

	query, v, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "meeting_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().GetContext(ctx, &id, q, v...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) Get(ctx context.Context, id int64) (*model.Meeting, error) {
	builder := sq.Select("id, topic_id, status, start_date, speaker_id, listener_id, created_at").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{"id": id}).
		Limit(1)

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "meeting_repository.Get",
		QueryRaw: query,
	}

	var res model.Meeting
	err = r.db.DB().GetContext(ctx, &res, q, v...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *repository) GetList(ctx context.Context, filter *Query) ([]*model.Meeting, error) {
	builder := sq.Select("id, unit_id, topic_id, status, start_date, speaker_id, listener_id, created_at").
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
		Name:     "meeting_repository.GetList",
		QueryRaw: query,
	}

	var res []*model.Meeting
	err = r.db.DB().SelectContext(ctx, &res, q, v...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *repository) Update(ctx context.Context, ids []int64, updateMeeting *model.UpdateMeeting) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": ids})

	if updateMeeting.UnitID.Valid {
		builder = builder.Set("unit_id", updateMeeting.UnitID.Int64)
	}
	if updateMeeting.TopicID.Valid {
		builder = builder.Set("topic_id", updateMeeting.TopicID.Int64)
	}
	if updateMeeting.Status.Valid {
		builder = builder.Set("status", updateMeeting.Status)
	}
	if updateMeeting.StartDate.Valid {
		builder = builder.Set("start_date", updateMeeting.StartDate.Time)
	}
	if updateMeeting.SpeakerID.Valid {
		builder = builder.Set("speaker_id", updateMeeting.SpeakerID.Int64)
	}
	if updateMeeting.ListenerID.Valid {
		builder = builder.Set("listener_id", updateMeeting.ListenerID.Int64)
	}

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "meeting_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, v...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetSpeakerCountByTopic(ctx context.Context, topicID int64, speakerID int64) (int64, error) {
	builder := sq.Select("count(*)").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{"topic_id": topicID}).
		Where(sq.Eq{"speaker_id": speakerID}).
		Where(sq.Eq{"status": model.MeetingStatusFinished})

	query, v, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "meeting_repository.GetSpeakerCountByTopic",
		QueryRaw: query,
	}

	var count int64
	err = r.db.DB().QueryRowContext(ctx, q, v...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetSpeakersStats ...
// select topic_id, speaker_id, count(*) from meeting group by topic_id, speaker_id, status having status='finished';
func (r *repository) GetSpeakersStats(ctx context.Context, unitID int64, topicID int64, excludeSpeakerID int64) ([]*model.Stats, error) {
	builder := sq.Select("unit_id, topic_id, speaker_id, count(*)").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{"topic_id": topicID}).
		Where(sq.Eq{"unit_id": unitID}).
		GroupBy("unit_id, topic_id, speaker_id, status").
		Having(sq.And{
			sq.Eq{"status": model.MeetingStatusFinished},
			sq.NotEq{"speaker_id": excludeSpeakerID},
		})

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "meeting_repository.GetSpeakersStats",
		QueryRaw: query,
	}

	var res []*model.Stats
	err = r.db.DB().SelectContext(ctx, &res, q, v...)
	if err != nil {
		return nil, err
	}

	return res, nil
}
