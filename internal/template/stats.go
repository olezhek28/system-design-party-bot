package template

const SpeakerStats = `🎯 Тема "{{ .TopicName }}" была пересказана {{ .Count }} раз(а).
`

const SpeakerStatsIntroduction = `А вот и статистика спикера {{ .FirstName }} {{ .LastName }} по разным темам.
🍺🏀🎖🥁🎸
`

const TopicStatsHeader = `📮 Тему "{{ .TopicName }}" пересказывали:
`

const TopicStats = `						🧩 {{ .FirstName }} {{ .LastName }} {{ .Count }} раз(а).
`

const TopicStatsIntroduction = `Статистика спикеров по разным темам.
🍺🏀🎖🥁🎸
`
