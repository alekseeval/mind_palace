package model

const (
	// Internal errors
	InternalServerError    = 1000
	WrongRequestParameters = 1001

	// User error codes
	UserNameUsed        = 2001
	UserTgIdUsed        = 2002
	UserNameTooLong     = 2003
	NoSuchUser          = 2004
	UserThemeLinkExists = 2005

	// Theme error codes
	ThemeExists           = 2006
	NoSuchMainTheme       = 2007
	MainThemeUserMismatch = 2008
	ThemeMainItself       = 2009
	NoteThemeLinkExists   = 2010
	ThemeHaveSubThemes    = 2011
	NoSuchTheme           = 2012

	DbError    = 3001
	UserExists = 3003
	TgIdInUse  = 3004
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

	InternalServerError: "Internal server error",
	DbError:             "DB error",
	UserExists:          "User already exists",
	TgIdInUse:           "Telegram id is used by other user",
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
