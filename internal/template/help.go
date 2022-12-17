package template

import "github.com/olezhek28/system-design-party-bot/internal/model/command"

const HelpMsg = `Справка по командам:
/` + command.GetCalendar + ` — посмотреть список твоих предстоящих встреч. 
/` + command.SetTimezone + ` — установись часовой пояс. Пример: /` + command.SetTimezone + ` 3, где 3 — это UTC+3.
/` + command.GetTimezone + ` — посмотреть установленный часовой пояс.
`
