package model

const (
	InternalServerError = 1000
	DbError             = 1001
	NoSuchUser          = 1002
	UserExists          = 1003
	TgIdInUse           = 1004
	NoSuchTheme         = 1005
	ThemeExists         = 1006
	NoSuchMainTheme     = 1007
)

var ErrMap = map[int]string{
	InternalServerError: "Internal server error",
	DbError:             "DB error",
	NoSuchUser:          "No such user",
	UserExists:          "User already exists",
	TgIdInUse:           "Telegram id is used by other user",
	NoSuchTheme:         "No such theme",
	ThemeExists:         "Theme already exists",
	NoSuchMainTheme:     "No such main theme",
}

type HttpError struct {
	Code    int    `json:"code"`
	Key     string `json:"key,omitempty"`
	Message string `json:"msg,omitempty"`
}

func NewHTTPError(code int, err error) *HttpError {
	he := &HttpError{
		Code:    code,
		Message: ErrMap[code],
	}
	if err != nil {
		he.Key = err.Error()
	}
	return he
}

func (e *HttpError) Error() string {
	return e.Key + ": " + e.Message
}
