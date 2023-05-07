package model

type IDAO interface {
	IUserDAO
	IThemeDAO
	INoteDAO
}

type IUserDAO interface {
	SaveUser(user User) (*User, error)
	GetUserByTgId(telegramIid int64) (*User, error)
	ChangeUser(user *User) (*User, error)
	DeleteUser(userId int) (int, error)
}

type IThemeDAO interface {
	CreateTheme(theme Theme) (*Theme, error)
	GetAllUserThemes(userId int) ([]*Theme, error)
	ChangeTheme(theme *Theme) (*Theme, error)
	DeleteTheme(themeId int) (int, error)
}

type INoteDAO interface {
	CreateNote(note Note) (*Note, error)
	GetAllUserNotesByTheme(userId int, themeId int) ([]*Note, error)
	ChangeNote(note *Note) (*Note, error)
	DeleteNote(noteId int) (int, error)
}
