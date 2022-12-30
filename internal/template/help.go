package template

import "github.com/olezhek28/system-design-party-bot/internal/model/command"

const HelpMsg = `Справка по командам:
/` + command.GetCalendar + ` — посмотреть список твоих предстоящих встреч. 
/` + command.SetTimezone + ` — установись часовой пояс. Пример: /` + command.SetTimezone + ` <b>3</b>, где <b>3</b> — это <b>UTC+3</b>. Если ты в <b>UTC-3</b>, то введи <b>-3</b>. А если вдруг твой часовой пояс имеет не целое количество часов, то введи часы и минуты через двоеточие, например, <b>3:30</b>, где <b>3:30</b> — это <b>UTC+3:30</b>.
/` + command.GetTimezone + ` — посмотреть установленный часовой пояс.
`
