package template

import (
	"github.com/olezhek28/system-design-party-bot/internal/model"
)

const NotificationAfterCreate = `{{ .Emoji }} Пользователь <b>{{ .FirstName }} {{ .LastName }}</b> (@{{ .TelegramUsername }}) запланировал с тобой встречу по теме <b>"{{ .TopicName}}"</b> из раздела <b>"{{ .UnitName}}"</b> на <b>{{ .StartDate }}</b> (в твоём часовом поясе).
`

const NotificationBeforeStart = `{{ .Emoji }} Йоу! Напоминаю, что через <b>` + model.ReminderTime + `</b> минут у тебя запланирована встреча с пользователем <b>{{ .FirstName }} {{ .LastName }}</b> (@{{ .TelegramUsername }}) по теме <b>"{{ .TopicName}}"</b> из раздела <b>"{{ .UnitName}}"</b>. Не забудь!
`
