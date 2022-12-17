package template

const SocialConnectionDescription = `{{ .Emoji }} Вот такие социальные связи у нас уже сформировались:
`

const SocialConnectionStudentName = `{{ .Emoji }} {{ .StudentFirstName }} {{ .StudentLastName }} (@{{ .StudentTelegramUsername }}) уже знаком(а) с:
`

const SocialConnection = `					🟢 {{ .PartnerFirstName }} {{ .PartnerLastName }} (@{{ .PartnerTelegramUsername }})
`

const SocialNotConnection = `					🔴 {{ .PartnerFirstName }} {{ .PartnerLastName }} (@{{ .PartnerTelegramUsername }})
`
