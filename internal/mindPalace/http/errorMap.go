package http

import (
	"MindPalace/internal/mindPalace/model"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func MapDBError(dbErr *pq.Error) *model.ServerError {
	var serverError *model.ServerError
	if dbErr.Code == "23505" { // unique constrain DB error
		switch dbErr.Constraint {
		case "users_name_key":
			serverError = model.NewServerError(model.UserNameUsed, dbErr)
		case "users_tg_id_key":
			serverError = model.NewServerError(model.UserTgIdUsed, dbErr)
		case "themes_title_user_id_key":
			serverError = model.NewServerError(model.ThemeExists, dbErr)
		}
	}

	if dbErr.Code == "23503" { // foreign key violation
		switch dbErr.Constraint {
		case "themes_user_id_fkey":
			serverError = model.NewServerError(model.UserThemeLinkExists, dbErr)
		case "themes_main_theme_id_fkey":
			serverError = model.NewServerError(model.NoSuchMainTheme, dbErr)
		case "notes_theme_id_fkey":
			serverError = model.NewServerError(model.NoteThemeLinkExists, dbErr)
		}
	}

	switch dbErr.Code { // Own codes
	case "80001":
		serverError = model.NewServerError(model.UserNameTooLong, dbErr)
	case "80002":
		serverError = model.NewServerError(model.NoSuchUser, dbErr)
	case "80003":
		serverError = model.NewServerError(model.MainThemeUserMismatch, dbErr)
	case "80004":
		serverError = model.NewServerError(model.NoSuchMainTheme, dbErr)
	case "80005":
		serverError = model.NewServerError(model.ThemeMainItself, dbErr)
	case "80006":
		serverError = model.NewServerError(model.ThemeHaveSubThemes, dbErr)
	case "80007":
		serverError = model.NewServerError(model.NoSuchTheme, dbErr)
	}

	if serverError == nil { // Unexpected DB error
		serverError = model.NewServerError(model.DbError, dbErr)
		log.Debug(*dbErr)
	}
	return serverError
}
