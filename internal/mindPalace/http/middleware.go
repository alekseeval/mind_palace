package http

import (
	"MindPalace/internal/mindPalace/model"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
)

func (s *HttpServer) requestLogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Read the content
		var bodyBytes []byte
		var parsedBody interface{}
		if ctx.Request().Body != nil {
			bodyBytes, _ = io.ReadAll(ctx.Request().Body)
			err := json.Unmarshal(bodyBytes, &parsedBody)
			if err != nil {
				parsedBody = nil
			}
		}
		// Restore the io.ReadCloser to its original state
		ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		s.logEntry.WithFields(log.Fields{
			"headers": ctx.Request().Header,
			"body":    parsedBody,
			"uri":     ctx.Request().URL.Path,
		}).Info("Request handled")
		return next(ctx)
	}
}

func (s *HttpServer) customHTTPErrorHandler(returnedErr error, c echo.Context) {
	s.logEntry.Error(returnedErr)
	if c.Response().Committed {
		return
	}

	// Catch db errors
	if dbErr, ok := returnedErr.(*pq.Error); ok {
		serverErr := MapDBError(dbErr)
		err := c.JSON(http.StatusInternalServerError, serverErr)
		if err != nil {
			s.logEntry.Error("Error occurred when returning HTTP error")
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
