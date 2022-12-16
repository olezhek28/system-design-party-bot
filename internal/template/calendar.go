package template

const CalendarDescription = `🍟 Список запланированных встреч, которые ещё не успели состояться.
`

const CalendarMeeting = `🍺 {{ .SpeakerFirstName }} {{ .SpeakerLastName }} (@{{ .SpeakerTelegramUsername }}) X {{ .ListenerFirstName }} {{ .ListenerLastName }} (@{{ .ListenerTelegramUsername }})
"{{ .TopicName }}" — {{ .StartDate }}.
`
