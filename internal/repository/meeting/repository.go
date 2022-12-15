package meeting_repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i Repository -o ./mocks/ -s "_minimock.go"

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/olezhek28/system-design-party-bot/internal/model"
	"github.com/olezhek28/system-design-party-bot/internal/pkg/db"
)

const (
	meetingTable = "meeting"
	studentTable = "student"
	topicTable   = "topic"
)

// Repository ...
type Repository interface {
	GetSpeakersStats(ctx context.Context, topicID int64, excludeSpeakerTelegramID int64) ([]*model.Stats, error)
	GetSuccessMeetingBySpeaker(ctx context.Context, speakerID int64) ([]*model.Meeting, error)
	GetSuccessMeeting(ctx context.Context) ([]*model.Meeting, error)
	CreateMeeting(ctx context.Context, meeting *model.Meeting) (int64, error)
	UpdateMeetingsStatus(ctx context.Context, status string, meetingIDs []int64) error
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
func (r *repository) GetSpeakersStats(ctx context.Context, topicID int64, excludeSpeakerTelegramID int64) ([]*model.Stats, error) {
	builder := sq.Select("t.name, s.id, s.first_name, s.last_name, s.telegram_username, count(*) ").
		PlaceholderFormat(sq.Dollar).
		From(meetingTable + " m").
		Join(studentTable + " s on m.speaker_id=s.id").
		Join(topicTable + " t on m.topic_id=t.id").
		Where(sq.Eq{"m.topic_id": topicID}).
		GroupBy("t.name, s.id, s.first_name, s.last_name, s.telegram_username, m.status").
		Having(sq.And{
			sq.Eq{"m.status": model.MeetingStatusFinished},
			sq.NotEq{"s.telegram_id": excludeSpeakerTelegramID},
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

func (r *repository) GetSuccessMeetingBySpeaker(ctx context.Context, speakerID int64) ([]*model.Meeting, error) {
	builder := sq.Select("id, topic_id, status, start_date, speaker_id, listener_id, created_at").
		PlaceholderFormat(sq.Dollar).
		From(meetingTable).
		Where(sq.Eq{"speaker_id": speakerID}).
		Where(sq.Eq{"status": model.MeetingStatusFinished})

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "meeting_repository.GetSuccessMeetingBySpeaker",
		QueryRaw: query,
	}

	var res []*model.Meeting
	err = r.db.DB().SelectContext(ctx, &res, q, v...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *repository) GetSuccessMeeting(ctx context.Context) ([]*model.Meeting, error) {
	builder := sq.Select("id, topic_id, status, start_date, speaker_id, listener_id, created_at").
		PlaceholderFormat(sq.Dollar).
		From(meetingTable).
		Where(sq.Eq{"status": model.MeetingStatusFinished})

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "meeting_repository.GetSuccessMeetingByTopic",
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
	builder := sq.Insert(meetingTable).
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
	builder := sq.Update(meetingTable).
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
