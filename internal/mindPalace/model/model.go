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

func (ua *UserAttributes) NewUserWithAttr() *User {
	user := User{
		Name:       ua.Name,
		TelegramId: ua.TelegramId,
	}
	return &user
}

type ThemeAttributes struct {
	Title       *string `json:"title"`
	MainThemeId *int    `json:"main_theme_id"`
}

func (ta *ThemeAttributes) NewThemeWithAttributes() *Theme {
	theme := Theme{
		Title:       ta.Title,
		MainThemeId: ta.MainThemeId,
	}
	return &theme
}

type NoteAttributes struct {
	Title      *string   `json:"title"`
	Text       *string   `json:"text"`
	NoteTypeId *NoteType `json:"note_type"`
	ThemeId    *int      `json:"theme_id"`
}

func (na *NoteAttributes) NewNoteWithAttributes() *Note {
	note := Note{
		Title:      na.Title,
		Text:       na.Text,
		NoteTypeId: na.NoteTypeId,
		ThemeId:    na.ThemeId,
	}
	return &note
}
