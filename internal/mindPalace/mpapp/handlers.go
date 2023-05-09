package mpapp

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// e.POST("/authorize")
func (s *HttpServer) authorize(c echo.Context) error {
	tgId, err := strconv.ParseInt(c.FormValue("tg_id"), 10, 64)
	if err != nil {
		return err
	}
	user, err := s.storage.GetUserByTgId(tgId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

// TODO: переписать
// e.POST("/register")
func (s *HttpServer) register(c echo.Context) error {

	tgId, err := strconv.ParseInt(c.FormValue("tg_id"), 10, 64)
	if err != nil {
		return err
	}
	user, err := s.storage.GetUserByTgId(tgId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
