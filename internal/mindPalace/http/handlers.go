package http

import (
	"MindPalace/internal/mindPalace/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// e.POST("/users")
func (s *HttpServer) createUser(c echo.Context) error {
	u := model.UserUpdate{}
	err := (&echo.DefaultBinder{}).BindBody(c, &u)
	if err != nil {
		return err
	}
	user := u.UpdateUser(&model.User{})
	dbUser, err := s.storage.SaveUser(*user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dbUser)
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
	return c.JSON(http.StatusOK, echo.Map{"id": deletedUserId})
}

// e.PATCH("/users/:id")
func (s *HttpServer) changeUser(c echo.Context) error {
	// read request parameters
	updatedUserParams := model.UserUpdate{}
	err := (&echo.DefaultBinder{}).BindBody(c, &updatedUserParams)
	if err != nil {
		return err
	}
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// update user in db
	user, err := s.storage.GetUserById(userId)
	if err != nil {
		return err
	}
	user = updatedUserParams.UpdateUser(user)
	dbUser, err := s.storage.ChangeUser(user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dbUser)
}
