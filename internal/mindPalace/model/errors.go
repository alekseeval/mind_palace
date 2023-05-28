package model

// Internal errors
const (
	InternalServerError = iota + 1000
	WrongRequestParameters
	DbError
)

const (
	// User error codes
	UserNameUsed = iota + 2000
	UserTgIdUsed
	UserNameTooLong
	NoSuchUser
	UserThemeLinkExists

	// Theme error codes
	ThemeExists
	NoSuchMainTheme
	MainThemeUserMismatch
	ThemeMainItself
	NoteThemeLinkExists
	ThemeHaveSubThemes
	NoSuchTheme

	// Note error codes
	NoteNoThemeProvided
	NoteNoTitleProvided
	NoteNoTextProvided
	NoSuchNote
	InvalidNoteType
)

var ErrMap = map[int]string{
	WrongRequestParameters: "Wrong request parameters",

	// Users
	UserNameUsed:        "User with this name already exists",
	UserTgIdUsed:        "User with this tg_id already exists",
	UserNameTooLong:     "User name should be at most 50 characters",
	NoSuchUser:          "No such user",
	UserThemeLinkExists: "For user exists non deleted theme",

	// Themes
	ThemeExists:           "Theme already exists",
	NoSuchMainTheme:       "No such main theme",
	MainThemeUserMismatch: "Main theme linked to other user",
	NoteThemeLinkExists:   "Theme linked to some notes",
	ThemeMainItself:       "Theme cant be main for itself",
	ThemeHaveSubThemes:    "Theme have sub themes",
	NoSuchTheme:           "No such theme",

	// Notes
	NoteNoThemeProvided: "Note must be linked to theme",
	NoteNoTitleProvided: "Note must have the Title",
	NoteNoTextProvided:  "Note must have the Text",
	NoSuchNote:          "No such note",
	InvalidNoteType:     "Invalid note type provided",

	InternalServerError: "Internal server error",
	DbError:             "DB error",
}

type ServerError struct {
	Code    int    `json:"code"`
	Message string `json:"msg,omitempty"`
	Detail  string `json:"detail,omitempty"`
}

func NewServerError(code int, err error) *ServerError {
	he := &ServerError{
		Code:    code,
		Message: ErrMap[code],
	}
	if err != nil {
		he.Detail = err.Error()
	}
	return he
}

func (e *ServerError) Error() string {
	return e.Detail + ": " + e.Message
}
