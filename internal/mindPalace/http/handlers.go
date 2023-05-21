package http

import (
	"MindPalace/internal/mindPalace/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// ---------------------------------------------------------------------------------------------------------------------
//  User API
// ---------------------------------------------------------------------------------------------------------------------
// e.POST("/users")
func (s *HttpServer) createUser(c echo.Context) error {
	userData := new(model.UserUpdate)
	err := c.Bind(&userData)
	if err != nil {
		return err
	}
	user := userData.UpdateUser(&model.User{})
	dbUser, err := s.storage.SaveUser(*user)
	if err != nil {
		return model.NewServerError(model.DbError, err)
	}
	return c.JSON(http.StatusCreated, dbUser)
}

// e.GET("/users/:id")
func (s *HttpServer) getUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	user, err := s.storage.GetUserById(userId)
	if err != nil {
		return model.NewServerError(model.DbError, err)
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
		return model.NewServerError(model.DbError, err)
	}
	return c.JSON(http.StatusOK, user)
}

// e.DELETE("/users/:id")
func (s *HttpServer) deleteUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = s.storage.DeleteUser(userId)
	if err != nil {
		return model.NewServerError(model.DbError, err)
	}
	return c.NoContent(http.StatusOK)
}

// e.PATCH("/users/:id")
func (s *HttpServer) editUser(c echo.Context) error {
	// read request parameters
	updatedUserParams := new(model.UserUpdate)
	err := c.Bind(&updatedUserParams)
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
		return model.NewServerError(model.DbError, err)
	}
	user = updatedUserParams.UpdateUser(user)
	dbUser, err := s.storage.ChangeUser(user)
	if err != nil {
		return model.NewServerError(model.DbError, err)
	}
	return c.JSON(http.StatusOK, dbUser)
}

// ---------------------------------------------------------------------------------------------------------------------
//  Themes API
// ---------------------------------------------------------------------------------------------------------------------
// e.POST("/themes")
func (s *HttpServer) createTheme(c echo.Context) error {
	themeData := new(model.ThemeUpdate)
	err := c.Bind(&themeData)
	if err != nil {
		return model.NewServerError(model.InternalServerError, err)
	}
	userIdHeader := c.Request().Header["Metadata-User"]
	var userName string
	if len(userIdHeader) != 0 {
		userName = userIdHeader[0]
	} else {
		userName = model.SystemUser
	}
	theme := themeData.UpdateTheme(&model.Theme{UserName: userName})
	theme, err = s.storage.SaveTheme(*theme)
	if err != nil {
		return model.NewServerError(model.DbError, err)
	}
	return c.JSON(http.StatusCreated, theme)
}

// e.GET("/themes")
func (s *HttpServer) getUserThemes(c echo.Context) error {
	userIdHeader := c.Request().Header["Metadata-User"]
	var userName *string
	if len(userIdHeader) != 0 {
		userName = &userIdHeader[0]
	}
	userThemes, err := s.storage.GetAllUserThemes(userName)
	if err != nil {
		return model.NewServerError(model.DbError, err)
	}
	return c.JSON(http.StatusOK, userThemes)
}

// e.DELETE("/themes/:id")
func (s *HttpServer) deleteTheme(c echo.Context) error {
	themeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return model.NewServerError(model.InternalServerError, err)
	}
	err = s.storage.DeleteTheme(themeId)
	if err != nil {
		return model.NewServerError(model.DbError, err)
	}
	return c.NoContent(http.StatusOK)
}

// e.PATCH("/theme/:theme_id")
func (s *HttpServer) editTheme(c echo.Context) error {
	themeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return model.NewServerError(model.InternalServerError, err)
	}
	themeData := new(model.ThemeUpdate)
	err = c.Bind(&themeData)
	if err != nil {
		return model.NewServerError(model.InternalServerError, err)
	}
	theme := themeData.UpdateTheme(&model.Theme{Id: themeId})
	theme, err = s.storage.ChangeTheme(theme)
	if err != nil {
		return model.NewServerError(model.DbError, err)
	}
	return c.JSON(http.StatusOK, theme)
}

// ---------------------------------------------------------------------------------------------------------------------
//  Notes API
// ---------------------------------------------------------------------------------------------------------------------
// e.POST("/users/:user_id/theme/:theme_id/note")
func (s *HttpServer) createNote(c echo.Context) error {

	return c.NoContent(http.StatusCreated)
}
