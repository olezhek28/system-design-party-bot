package template

const SocialConnectionDescription = `{{ .Emoji }} Вот такие социальные связи у нас уже сформировались:
`

const SocialOwnConnectionDescription = `{{ .Emoji }} Ты уже знаком с этими людьми:
`

const SocialConnectionStudentName = `{{ .Emoji }} {{ .StudentFirstName }} {{ .StudentLastName }} (@{{ .StudentTelegramUsername }}) уже знаком(а) с этими людьми:
`

const SocialConnection = `					🟢 {{ .PartnerFirstName }} {{ .PartnerLastName }} (@{{ .PartnerTelegramUsername }})
`

const SocialNotConnection = `					🔴 {{ .PartnerFirstName }} {{ .PartnerLastName }} (@{{ .PartnerTelegramUsername }})
`
