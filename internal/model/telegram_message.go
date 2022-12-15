package model

type User struct {
	ID        int64
	IsBot     bool
	FirstName string
	LastName  string
	UserName  string
}

type TelegramMessage struct {
	ID        int64
	From      *User
	Text      string
	Command   string
	Arguments []string
}
