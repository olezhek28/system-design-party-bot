package template

const SpeakerDescription = `Я нашел для тебя партнера для тренировки темы "{{ .TopicName }}".
Попробуй написать ему(ей) в личку, чтобы договориться о тренировке.
👤 {{ .FirstName }} {{ .LastName }}
✏️ {{ .TelegramUsername }}
🏆 {{ .Count }} раз пересказывал(а) тему "{{ .TopicName }}"
`
