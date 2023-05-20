package model

type IDAO interface {
	IUserDAO
	IThemeDAO
	INoteDAO
}

type IUserDAO interface {
	SaveUser(user User) (*User, error)
	GetUserByTgId(telegramId int64) (*User, error)
	GetUserByName(userId int) (*User, error)
	ChangeUser(user *User) (*User, error)
	DeleteUser(userId int) (int, error)
}

type IThemeDAO interface {
	SaveTheme(theme Theme) (*Theme, error)
	GetAllUserThemes(userName string) ([]*Theme, error)
	ChangeTheme(theme *Theme) (*Theme, error)
	DeleteTheme(themeId int) (int, error)
}

type INoteDAO interface {
	SaveNote(note Note) (*Note, error)
	GetAllNotesByTheme(themeId int) ([]*Note, error)
	ChangeNote(note *Note) (*Note, error)
	DeleteNote(noteId int) (int, error)
}
