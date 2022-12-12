package model

type Stats struct {
	SpeakerFirstName        string `db:"first_name"`
	SpeakerLastName         string `db:"last_name"`
	SpeakerTelegramNickname string `db:"telegram_username"`
	TopicID                 int64  `db:"topic_id"`
	Count                   int64  `db:"count"`
}
