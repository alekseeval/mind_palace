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
