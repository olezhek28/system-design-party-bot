package model

type Stats struct {
	TopicID   int64 `db:"topic_id"`
	SpeakerID int64 `db:"speaker_id"`
	Count     int64 `db:"count"`
}
