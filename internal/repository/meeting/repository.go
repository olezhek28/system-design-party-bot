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
	GetSpeakersStats(ctx context.Context, topicID int64, excludeSpeakerID int64) ([]*model.Stats, error)
	GetFinishedMeetingBySpeaker(ctx context.Context, speakerID int64) ([]*model.Meeting, error)
	GetMeetingsByStatus(ctx context.Context, status string) ([]*model.Meeting, error)
	CreateMeeting(ctx context.Context, meeting *model.Meeting) (int64, error)
	UpdateMeetingsStatus(ctx context.Context, status string, meetingIDs []int64) error
	GetSpeakerCountByTopic(ctx context.Context, topicID int64, speakerID int64) (int64, error)
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

// GetSpeakersStats ...
// select topic_id, speaker_id, count(*) from meeting group by topic_id, speaker_id, status having status='finished';
func (r *repository) GetSpeakersStats(ctx context.Context, topicID int64, excludeSpeakerID int64) ([]*model.Stats, error) {
	builder := sq.Select("topic_id, speaker_id, count(*)").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{"topic_id": topicID}).
		GroupBy("topic_id, speaker_id, status").
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

func (r *repository) GetFinishedMeetingBySpeaker(ctx context.Context, speakerID int64) ([]*model.Meeting, error) {
	builder := sq.Select("id, topic_id, status, start_date, speaker_id, listener_id, created_at").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{"speaker_id": speakerID}).
		Where(sq.Eq{"status": model.MeetingStatusFinished})

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "meeting_repository.GetFinishedMeetingBySpeaker",
		QueryRaw: query,
	}

	var res []*model.Meeting
	err = r.db.DB().SelectContext(ctx, &res, q, v...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *repository) GetMeetingsByStatus(ctx context.Context, status string) ([]*model.Meeting, error) {
	builder := sq.Select("id, topic_id, status, start_date, speaker_id, listener_id, created_at").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{"status": status})

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "meeting_repository.GetMeetingsByStatus",
		QueryRaw: query,
	}

	var res []*model.Meeting
	err = r.db.DB().SelectContext(ctx, &res, q, v...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *repository) CreateMeeting(ctx context.Context, meeting *model.Meeting) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("topic_id", "status", "start_date", "speaker_id", "listener_id").
		Values(meeting.TopicID, meeting.Status, meeting.StartDate, meeting.SpeakerID, meeting.ListenerID).
		Suffix("RETURNING id")

	query, v, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "meeting_repository.CreateMeeting",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().GetContext(ctx, &id, q, v...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) UpdateMeetingsStatus(ctx context.Context, status string, meetingIDs []int64) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set("status", status).
		Where(sq.Eq{"id": meetingIDs})

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "meeting_repository.UpdateMeetingStatus",
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
