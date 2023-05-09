package mpapp

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// e.POST("/authorize")
func (s *HttpServer) authorize(c echo.Context) error {

	return c.String(http.StatusOK, "Hello world!")
}

func (s *HttpServer) register(c echo.Context) error {

	return c.String(http.StatusOK, "Hello world!")
}
