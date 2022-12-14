package model

import "time"

type Meeting struct {
	ID         int64     `db:"id"`
	TopicID    int64     `db:"topic_id"`
	Status     string    `db:"status"`
	StartDate  time.Time `db:"start_date"`
	SpeakerID  int64     `db:"speaker_id"`
	ListenerID int64     `db:"listener_id"`
	CreatedAt  time.Time `db:"created_at"`
}
