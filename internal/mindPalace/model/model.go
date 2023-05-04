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
	Id          int    `db:"id"`
	Title       string `db:"title"`
	MainThemeId *int   `db:"main_theme_id"`
	UserId      *int   `db:"user_id"`
}

type Note struct {
	Id         int      `db:"id"`
	Title      string   `db:"title"`
	Text       string   `db:"text"`
	NoteTypeId NoteType `db:"note_type"`
	ThemeId    int      `db:"theme_id"`
	UserId     int      `db:"user_id"`
}
