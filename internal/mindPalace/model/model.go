package model

type NoteType int

const (
	SimpleNote NoteType = iota + 1
	Question
	Task
)

type User struct {
	Id         int     `db:"id" json:"id"`
	Name       *string `db:"name" json:"name"`
	TelegramId *int64  `db:"tg_id" json:"tg_id"`
}

type UserUpdate struct {
	Name       *string `db:"name" json:"name"`
	TelegramId *int64  `db:"tg_id" json:"tg_id"`
}

func (u *UserUpdate) UpdateUser(user *User) *User {
	if u.Name != nil {
		user.Name = u.Name
	}
	if u.TelegramId != nil {
		user.TelegramId = u.TelegramId
	}
	return user
}

type Theme struct {
	Id          int    `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	MainThemeId *int   `db:"main_theme_id" json:"main_theme_id"`
	UserId      *int   `db:"user_id" json:"user_id"`
}

type Note struct {
	Id         int      `db:"id" json:"id"`
	Title      string   `db:"title" json:"title"`
	Text       string   `db:"text" json:"text"`
	NoteTypeId NoteType `db:"note_type" json:"note_type"`
	ThemeId    int      `db:"theme_id" json:"theme_id"`
	UserId     int      `db:"user_id" json:"user_id"`
}
