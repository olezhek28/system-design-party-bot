package template

const SpeakerStatsUnitIntroduction = `					🚧 Раздел <b>"{{ .UnitName }}"</b>
`

const SpeakerStats = `									🎯 Тема "{{ .TopicName }}" была пересказана — <b>{{ .Count }} раз(а)</b>.
`

const SpeakerStatsIntroduction = `🤹‍♂️ А вот и статистика спикера <b>{{ .FirstName }} {{ .LastName }}</b> по разным темам.
`

const TopicStatsHeader = `📮 Тему "{{ .TopicName }}" пересказывали:
`

const TopicStats = `						🧩 {{ .FirstName }} {{ .LastName }} (@{{ .TelegramUsername }}) — {{ .Count }} раз(а).
`

const TopicStatsIntroduction = `Статистика спикеров по разным темам.
`
