package mpapp

type DAO interface {
	UserDAO
	ThemeDAO
	NoteDAO
}

type UserDAO interface {
	createUser()
	getUser()
	changeUser()
	deleteUser()
}

type ThemeDAO interface {
	createTheme()
	getTheme()
	changeTheme()
	deleteTheme()
}

type NoteDAO interface {
	createNote()
	getNote()
	changeNote()
	deleteNote()
}
