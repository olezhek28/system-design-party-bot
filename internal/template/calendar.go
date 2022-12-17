package template

const CalendarDescription = `{{ .Emoji }} Список запланированных встреч, которые ещё не успели состояться.
`

const CalendarMeeting = `{{ .Emoji }} {{ .SpeakerFirstName }} {{ .SpeakerLastName }} (@{{ .SpeakerTelegramUsername }}) X {{ .ListenerFirstName }} {{ .ListenerLastName }} (@{{ .ListenerTelegramUsername }})
"{{ .TopicName }}" — {{ .StartDate }}.
`
