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

func customHTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	// Catch db errors
	if dbErr, ok := err.(*pq.Error); ok {
		if dbErr.Code == "23505" {
			switch dbErr.Constraint {
			case "users_name_key":
				err = c.JSON(http.StatusInternalServerError, model.NewServerError(model.UserNameUsed, dbErr))
				if err != nil {
					log.Error(err)
				}
				return
			case "users_tg_id_key":
				err = c.JSON(http.StatusInternalServerError, model.NewServerError(model.UserTgIdUsed, dbErr))
				if err != nil {
					log.Error(err)
				}
				return
			}
		}
	}

	he, ok := err.(*model.ServerError)
	if ok {
		err = c.JSON(http.StatusInternalServerError, he)
		if err != nil {
			log.Error(err)
		}
	} else {
		err = c.JSON(http.StatusInternalServerError, model.NewServerError(model.InternalServerError, err))
		if err != nil {
			log.Error(err)
		}
	}
}
