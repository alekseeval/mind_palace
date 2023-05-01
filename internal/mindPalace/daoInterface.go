package mindPalace

type DAO interface {
	UserDAO
	ThemeDAO
	NoteDAO
}

type UserDAO interface {
	SaveUser(user User) (*User, error)
	GetUserByTgId(telegramIid int64) (*User, error)
	ChangeUser(user *User) (*User, error)
	DeleteUser(id int) (*User, error)
}

type ThemeDAO interface {
	CreateTheme(theme Theme) (*Theme, error)
	GetAllUserThemes(user *User) ([]*Theme, error)
	ChangeTheme(theme *Theme) (*Theme, error)
	DeleteTheme(id int) (*Theme, error)
}

type NoteDAO interface {
	CreateNote(note Note) (*Note, error)
	GetAllUserNotesByTheme(user *User, theme *Theme) ([]*Note, error)
	ChangeNote(note *Note) (*Note, error)
	DeleteNote(id int) (*Note, error)
}
