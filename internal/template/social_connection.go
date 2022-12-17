package template

const SocialConnectionDescription = `🍟 Вот такие социальные связи у нас уже сформировались:
`

const SocialConnectionStudentName = `🦊 {{ .StudentFirstName }} {{ .StudentLastName }} уже знаком(а) с:
`

const SocialConnection = `					🟢 {{ .PartnerFirstName }} {{ .PartnerLastName }}
`

const SocialNotConnection = `					🔴 {{ .PartnerFirstName }} {{ .PartnerLastName }}
`
