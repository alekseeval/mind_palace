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
		return model.NewServerError(model.WrongRequestParameters, err)
	}
	user := userData.UpdateUser(&model.User{})
	dbUser, err := s.storage.SaveUser(*user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, dbUser)
}

// GET("/users/:id")
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

// GET("/telegram/users/:tg_id")
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

// DELETE("/users/:id")
func (s *HttpServer) deleteUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = s.storage.DeleteUser(userId)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

// PATCH("/users/:id")
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
		return err
	}
	user = updatedUserParams.UpdateUser(user)
	dbUser, err := s.storage.ChangeUser(user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dbUser)
}

// ---------------------------------------------------------------------------------------------------------------------
//  Themes API
// ---------------------------------------------------------------------------------------------------------------------
// POST("/themes")
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
	theme := themeData.UpdateTheme(&model.Theme{UserName: &userName})
	theme, err = s.storage.SaveTheme(*theme)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, theme)
}

// GET("/themes")
func (s *HttpServer) getUserThemes(c echo.Context) error {
	userIdHeader := c.Request().Header["Metadata-User"]
	var userName *string
	if len(userIdHeader) != 0 {
		userName = &userIdHeader[0]
	}
	userThemes, err := s.storage.GetAllUserThemes(userName)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, userThemes)
}

// DELETE("/themes/:id")
func (s *HttpServer) deleteTheme(c echo.Context) error {
	themeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return model.NewServerError(model.InternalServerError, err)
	}
	err = s.storage.DeleteTheme(themeId)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

// PATCH("/theme/:theme_id")
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
		return err
	}
	return c.JSON(http.StatusOK, theme)
}

// ---------------------------------------------------------------------------------------------------------------------
//  Notes API
// ---------------------------------------------------------------------------------------------------------------------
// POST("/themes/:theme_id/notes")
func (s *HttpServer) createNote(c echo.Context) error {

	themeId, err := strconv.Atoi(c.Param("theme_id"))
	if err != nil {
		return model.NewServerError(model.InternalServerError, err)
	}
	var noteData model.NoteAttributes
	err = c.Bind(&noteData)
	if err != nil {
		return model.NewServerError(model.InternalServerError, err)
	}
	note := new(model.Note)
	note.ThemeId = &themeId
	note = noteData.UpdateNote(note)
	note, err = s.storage.SaveNote(*note)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, note)
}

// GET("/themes/:theme_id/notes")
func (s *HttpServer) getNotes(c echo.Context) error {
	themeId, err := strconv.Atoi(c.Param("theme_id"))
	if err != nil {
		return model.NewServerError(model.InternalServerError, err)
	}
	notes, err := s.storage.GetAllNotesByTheme(themeId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, notes)
}

// DELETE("/notes/:note_id")
func (s *HttpServer) deleteNote(c echo.Context) error {
	noteId, err := strconv.Atoi(c.Param("note_id"))
	if err != nil {
		return model.NewServerError(model.InternalServerError, err)
	}
	err = s.storage.DeleteNote(noteId)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

// PATCH("/notes/:note_id")
func (s *HttpServer) editNote(c echo.Context) error {
	noteId, err := strconv.Atoi(c.Param("note_id"))
	if err != nil {
		return model.NewServerError(model.InternalServerError, err)
	}
	var noteData model.NoteAttributes
	err = c.Bind(&noteData)
	if err != nil {
		return model.NewServerError(model.InternalServerError, err)
	}
	note := new(model.Note)
	note.Id = noteId
	note = noteData.UpdateNote(note)
	note, err = s.storage.ChangeNote(note)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, note)
}
