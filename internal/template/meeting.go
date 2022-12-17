package template

const CreateMeetingDescription = `{{ .Emoji }} Встреча создана на {{ .StartDate }}.
🎭 Твоим партнером будет {{ .FirstName }} {{ .LastName }}
🍾 Не забудь сообщить мне, когда закончишь тренировку, чтобы я мог обновить статистику. 
🥕 Если встреча не состоится, то отмени ее пожалуйста.
`

const FinishMeetingDescription = `🎉 Кайф! Поздравляю с успешно проведенной встречей! Статистику я обновил - движем дальше:)
`

const CancelMeetingDescription = `🔕 Встреча отменена. Можешь прям сейчас найти нового партнера для пересказа или потом. Главное не забудь;)
`

const MeetingConfirmation = `{{ .Emoji }} Назначить встречу на {{ .StartDate }}?
`
