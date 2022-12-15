package model

type Stats struct {
	SpeakerID               int64  `db:"id"`
	SpeakerFirstName        string `db:"first_name"`
	SpeakerLastName         string `db:"last_name"`
	SpeakerTelegramNickname string `db:"telegram_username"`
	TopicName               string `db:"name"`
	Count                   int64  `db:"count"`
}
