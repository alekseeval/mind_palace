package http

import (
	"MindPalace/internal/mindPalace/model"
	"github.com/lib/pq"
)

func MapDBError(dbErr *pq.Error) *model.ServerError {
	var serverError *model.ServerError
	if dbErr.Code == "23505" { // unique constrain DB error
		switch dbErr.Constraint {
		case "users_name_key":
			serverError = model.NewServerError(model.UserNameUsed, dbErr)
		case "users_tg_id_key":
			serverError = model.NewServerError(model.UserTgIdUsed, dbErr)
		}
	}

	if dbErr.Code == "23503" { // foreign key violation
		switch dbErr.Constraint {
		case "themes_user_id_fkey":
			serverError = model.NewServerError(model.UserThemeExists, dbErr)
		}
	}

	switch dbErr.Code { // Own codes
	case "80001":
		serverError = model.NewServerError(model.UserNameTooLong, dbErr)
	case "80002":
		serverError = model.NewServerError(model.NoSuchUser, dbErr)
	}

	if serverError == nil { // Unexpected DB error
		serverError = model.NewServerError(model.DbError, dbErr)
	}
	return serverError
}
