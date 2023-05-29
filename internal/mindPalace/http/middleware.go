package http

import (
	"MindPalace/internal/mindPalace/model"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"net/http"
)

func (s *HttpServer) logMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func (s *HttpServer) customHTTPErrorHandler(returnedErr error, c echo.Context) {
	if c.Response().Committed {
		s.logEntry.Error("Response was already committed when starting to handle error")
		return
	}

	// Catch db errors
	if dbErr, ok := returnedErr.(*pq.Error); ok {
		serverErr := MapDBError(dbErr)
		err := c.JSON(http.StatusInternalServerError, serverErr)
		if err != nil {
			s.logEntry.Error("Error occurred when return HTTP error")
		}
		return
	}

	// Catch error which already are a model.ServerError
	he, ok := returnedErr.(*model.ServerError)
	if ok {
		err := c.JSON(http.StatusInternalServerError, he)
		if err != nil {
			s.logEntry.Error(err)
		}
	} else { // When unexpected error occurred
		err := c.JSON(http.StatusInternalServerError, model.NewServerError(model.InternalServerError, returnedErr))
		if err != nil {
			s.logEntry.Error(err)
		}
	}
}
