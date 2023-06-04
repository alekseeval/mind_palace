package model

type NoteType int

const (
	SimpleNote NoteType = iota + 1
	Question
	Task

	SystemUser = "system"
)

type User struct {
	Id         int     `db:"id" json:"id"`
	Name       *string `db:"name" json:"name"`
	TelegramId *int64  `db:"tg_id" json:"tg_id"`
}

type Theme struct {
	Id          int     `db:"id" json:"id"`
	Title       *string `db:"title" json:"title"`
	MainThemeId *int    `db:"main_theme_id" json:"main_theme_id"`
	UserName    *string `db:"user_name" json:"user"`
}

type Note struct {
	Id         int       `db:"id" json:"id"`
	Title      *string   `db:"title" json:"title"`
	Text       *string   `db:"text" json:"text"`
	NoteTypeId *NoteType `db:"note_type" json:"note_type"`
	ThemeId    *int      `db:"theme_id" json:"theme_id"`
}

type UserAttributes struct {
	Name       *string `json:"name"`
	TelegramId *int64  `json:"tg_id"`
}

func (u *UserAttributes) UpdateUser(user *User) *User {
	user.Name = u.Name
	if u.TelegramId != nil {
		user.TelegramId = u.TelegramId
	}
	return user
}

type ThemeAttributes struct {
	Title       *string `json:"title"`
	MainThemeId *int    `json:"main_theme_id"`
}

func (tu *ThemeAttributes) UpdateTheme(theme *Theme) *Theme {
	if tu.Title != nil {
		theme.Title = tu.Title
	}
	if tu.MainThemeId != nil {
		theme.MainThemeId = tu.MainThemeId
	}
	return theme
}

type NoteAttributes struct {
	Title      *string   `json:"title"`
	Text       *string   `json:"text"`
	NoteTypeId *NoteType `json:"note_type"`
	ThemeId    *int      `json:"theme_id"`
}

func (nu NoteAttributes) UpdateNote(note *Note) *Note {
	if nu.Title != nil {
		note.Title = nu.Title
	}
	if nu.Text != nil {
		note.Text = nu.Text
	}
	if nu.NoteTypeId != nil {
		note.NoteTypeId = nu.NoteTypeId
	}
	if nu.ThemeId != nil {
		note.ThemeId = nu.ThemeId
	}
	return note
}
