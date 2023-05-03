package model

type NoteType int

const (
	SimpleNote NoteType = iota + 1
	Question
	Task
)

type User struct {
	Id         int     `db:"id"`
	Name       *string `db:"name"`
	TelegramId *int64  `db:"tg_id"`
}

type Theme struct {
	Id          int
	Title       string
	MainThemeId *int
	UserId      *int
}

type Note struct {
	Id         int
	Title      string
	Text       string
	NoteTypeId NoteType
	ThemeId    int
	UserId     int
}
