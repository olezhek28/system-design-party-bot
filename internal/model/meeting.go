package model

import (
	"database/sql"
	"time"
)

const (
	ReminderTime = "10"

	TimeFormat = "02-Jan-2006 15:04"
)

const (
	// MeetingStatusNew - new meeting
	MeetingStatusNew = "new"
	// MeetingStatusCanceled - canceled meeting
	MeetingStatusCanceled = "canceled"
	// MeetingStatusFinished - finished meeting
	MeetingStatusFinished = "finished"
)

type Meeting struct {
	ID         int64     `db:"id"`
	TopicID    int64     `db:"topic_id"`
	Status     string    `db:"status"`
	StartDate  time.Time `db:"start_date"`
	SpeakerID  int64     `db:"speaker_id"`
	ListenerID int64     `db:"listener_id"`
	CreatedAt  time.Time `db:"created_at"`
}

type UpdateMeeting struct {
	TopicID    sql.NullInt64
	Status     sql.NullString
	StartDate  sql.NullTime
	SpeakerID  sql.NullInt64
	ListenerID sql.NullInt64
}
