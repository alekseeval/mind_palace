package http

import (
	"MindPalace/internal/mindPalace/model"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func logMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func customHTTPErrorHandler(returnedErr error, c echo.Context) {
	if c.Response().Committed {
		log.Error("Response was already committed when starting to handle error")
		return
	}

	// Catch db errors
	if dbErr, ok := returnedErr.(*pq.Error); ok {
		serverErr := MapDBError(dbErr)
		err := c.JSON(http.StatusInternalServerError, serverErr)
		if err != nil {
			log.Error("Error occurred when return HTTP error")
		}
		return
	}

	// Catch error which already are a model.ServerError
	he, ok := returnedErr.(*model.ServerError)
	if ok {
		err := c.JSON(http.StatusInternalServerError, he)
		if err != nil {
			log.Error(err)
		}
	} else { // When unexpected error occurred
		err := c.JSON(http.StatusInternalServerError, model.NewServerError(model.InternalServerError, returnedErr))
		if err != nil {
			log.Error(err)
		}
	}
}

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
