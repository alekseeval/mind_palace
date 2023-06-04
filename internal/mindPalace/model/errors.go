package model

import "github.com/lib/pq"

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

var ErrDescriptionMap = map[int]string{
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
		Message: ErrDescriptionMap[code],
	}
	if err != nil {
		he.Detail = err.Error()
	}
	return he
}

func (e *ServerError) Error() string {
	return e.Detail + ": " + e.Message
}

func MapDBError(dbErr *pq.Error) *ServerError {
	var serverError *ServerError
	if dbErr.Code == "23505" { // unique constrain DB error
		switch dbErr.Constraint {
		case "users_name_key":
			serverError = NewServerError(UserNameUsed, dbErr)
		case "users_tg_id_key":
			serverError = NewServerError(UserTgIdUsed, dbErr)
		case "themes_title_user_id_key":
			serverError = NewServerError(ThemeExists, dbErr)
		}
	}

	if dbErr.Code == "23503" { // foreign key violation
		switch dbErr.Constraint {
		case "themes_user_id_fkey":
			serverError = NewServerError(UserThemeLinkExists, dbErr)
		case "themes_main_theme_id_fkey":
			serverError = NewServerError(NoSuchMainTheme, dbErr)
		case "notes_theme_id_fkey":
			serverError = NewServerError(NoteThemeLinkExists, dbErr)
		}
	}

	switch dbErr.Code { // Own codes
	case "80001":
		serverError = NewServerError(UserNameTooLong, dbErr)
	case "80002":
		serverError = NewServerError(NoSuchUser, dbErr)
	case "80003":
		serverError = NewServerError(MainThemeUserMismatch, dbErr)
	case "80004":
		serverError = NewServerError(NoSuchMainTheme, dbErr)
	case "80005":
		serverError = NewServerError(ThemeMainItself, dbErr)
	case "80006":
		serverError = NewServerError(ThemeHaveSubThemes, dbErr)
	case "80007":
		serverError = NewServerError(NoSuchTheme, dbErr)
	case "80008":
		serverError = NewServerError(NoteNoThemeProvided, dbErr)
	case "80009":
		serverError = NewServerError(NoteNoTitleProvided, dbErr)
	case "80011":
		serverError = NewServerError(InvalidNoteType, dbErr)
	case "80012":
		serverError = NewServerError(NoteNoTextProvided, dbErr)
	case "80013":
		serverError = NewServerError(NoSuchNote, dbErr)
	}

	if serverError == nil { // Unexpected DB error
		serverError = NewServerError(DbError, dbErr)
	}
	return serverError
}
