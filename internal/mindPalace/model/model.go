package model

type NoteType int

const (
	SimpleNote NoteType = iota + 1
	Question
	Task
)

type User struct {
	Id         int
	Name       string
	TelegramId int64
}

type Theme struct {
	Id          int
	Title       string
	MainThemeId int
	UserId      int
}

type Note struct {
	Id         int
	Title      string
	Text       string
	NoteTypeId NoteType
	ThemeId    int
	UserId     int
}
