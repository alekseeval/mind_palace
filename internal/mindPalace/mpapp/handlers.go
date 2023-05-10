package mpapp

import (
	"MindPalace/internal/mindPalace/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// e.POST("/users")
func (s *HttpServer) createUser(c echo.Context) error {
	u := model.User{}
	err := (&echo.DefaultBinder{}).BindBody(c, &u)
	if err != nil {
		return err
	}
	user, err := s.storage.SaveUser(u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// e.GET("/users/:id")
func (s *HttpServer) getUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	user, err := s.storage.GetUserById(userId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
