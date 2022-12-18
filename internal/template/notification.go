package template

import (
	"github.com/olezhek28/system-design-party-bot/internal/model"
)

const NotificationAfterCreate = `{{ .Emoji }} Пользователь <b>{{ .FirstName }} {{ .LastName }}</b> (@{{ .TelegramUsername }}) запланировал с тобой встречу по теме <b>"{{ .TopicName}}"</b> на <b>{{ .StartDate }}</b> (в твоём часов поясе).
`

const NotificationBeforeStart = `{{ .Emoji }} Йоу! Напоминаю, что через ` + string(model.ReminderTime) + ` у тебя запланирована встреча с пользователем <b>{{ .FirstName }} {{ .LastName }}</b> (@{{ .TelegramUsername }}) по теме <b>"{{ .TopicName}}"</b>. Не забудь!
`
