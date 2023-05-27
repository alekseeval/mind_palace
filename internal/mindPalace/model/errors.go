package model

const (
	// Internal errors
	InternalServerError    = 1000
	WrongRequestParameters = 1001

	// User error codes
	UserNameUsed    = 2001
	UserTgIdUsed    = 2002
	UserNameTooLong = 2003
	NoSuchUser      = 2004
	UserThemeExists = 2005

	DbError         = 3001
	UserExists      = 3003
	TgIdInUse       = 3004
	NoSuchTheme     = 3005
	ThemeExists     = 3006
	NoSuchMainTheme = 3007
)

var ErrMap = map[int]string{
	WrongRequestParameters: "Wrong request parameters",

	UserNameUsed:    "User with this name already exists",
	UserTgIdUsed:    "User with this tg_id already exists",
	UserNameTooLong: "User name should be at most 50 characters",
	NoSuchUser:      "No such user",
	UserThemeExists: "For user exists non deleted theme",

	InternalServerError: "Internal server error",
	DbError:             "DB error",
	UserExists:          "User already exists",
	TgIdInUse:           "Telegram id is used by other user",
	NoSuchTheme:         "No such theme",
	ThemeExists:         "Theme already exists",
	NoSuchMainTheme:     "No such main theme",
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
