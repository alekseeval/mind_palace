package model

const (
	InternalServerError = 1000

	// User error codes
	UserNameUsed    = 1001
	UserTgIdUsed    = 1002
	UserNameTooLong = 1003

	DbError         = 2001
	NoSuchUser      = 2002
	UserExists      = 2003
	TgIdInUse       = 2004
	NoSuchTheme     = 2005
	ThemeExists     = 2006
	NoSuchMainTheme = 2007
)

var ErrMap = map[int]string{
	UserNameUsed:    "User with this name already exists",
	UserTgIdUsed:    "User with this tg_id already exists",
	UserNameTooLong: "User name should be at most 50 characters",

	InternalServerError: "Internal server error",
	DbError:             "DB error",
	NoSuchUser:          "No such user",
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
