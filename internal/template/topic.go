package template

const TopicDescription = `🚩 №{{ .ID }}
📚 {{ .Name }}
📩 {{ .Description }}
`

const TopicGuidelines = `{{ .FirstName }}, если ты определился с темой, то самое время подобрать спикера,
который расскажет тебе эту тему. Для этого отправь мне команду /find_speaker topic_id, где topic_id - это 
номер темы, который указан напротив 🚩.
`
