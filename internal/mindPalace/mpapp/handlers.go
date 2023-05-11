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

// e.GET("/telegram/users/:tg_id")
func (s *HttpServer) getUserByTgId(c echo.Context) error {
	userTgId, err := strconv.ParseInt(c.Param("tg_id"), 10, 64)
	if err != nil {
		return err
	}
	user, err := s.storage.GetUserByTgId(userTgId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// e.DELETE("/users/:id")
func (s *HttpServer) deleteUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	deletedUserId, err := s.storage.DeleteUser(userId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, deletedUserId)
}

// e.PATCH("/users/:id")
func (s *HttpServer) changeUser(c echo.Context) error {
	u := model.User{}
	err := (&echo.DefaultBinder{}).BindBody(c, &u)
	if err != nil {
		return err
	}
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	u.Id = userId
	user, err := s.storage.ChangeUser(&u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
