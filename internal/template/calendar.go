package template

const CalendarDescription = `{{ .Emoji }} Список твоих предстоящих встреч:
`

const CalendarMeeting = `{{ .Emoji }} С пользователем {{ .ListenerFirstName }} {{ .ListenerLastName }} (@{{ .ListenerTelegramUsername }})
по теме "{{ .TopicName }}" — {{ .StartDate }}.
`

const CalendarAllDescription = `{{ .Emoji }} Список предстоящих встреч всех студентов:
`

const CalendarAllMeeting = `{{ .Emoji }} {{ .SpeakerFirstName }} {{ .SpeakerLastName }} (@{{ .SpeakerTelegramUsername }}) X {{ .ListenerFirstName }} {{ .ListenerLastName }} (@{{ .ListenerTelegramUsername }})
"{{ .TopicName }}" — {{ .StartDate }}.
`
