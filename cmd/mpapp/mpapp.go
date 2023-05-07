package main

import (
	"MindPalace/internal/mindPalace/configuration"
	"MindPalace/internal/mindPalace/dal"
	"MindPalace/internal/mindPalace/model"
	"fmt"
	"log"
)

var PATH_TO_CONFIG string = "/home/reserv/GolandProjects/MindPalace/internal/mindPalace/config.yaml"

func main() {
	config, err := configuration.ReadConfig(PATH_TO_CONFIG)
	if err != nil {
		log.Fatal("Error occurred when read config file\t", err)
	}

	dbDAO, err := dal.NewPostgresDB(config)
	if err != nil {
		log.Fatal("Failed to create connection to DB\t", err)
	}

	testUsers(dbDAO)
	testThemes(dbDAO)
	testNotes(dbDAO)
}

func testUsers(dbDAO model.DAO) {
	var userTgId int64 = 666
	var user *model.User

	// --------------------------------------------------------
	var userName *string = nil
	user, err := dbDAO.SaveUser(model.User{
		Name:       userName,
		TelegramId: &userTgId,
	})
	if err != nil {
		log.Fatal("Failed to create user\t", err)
	}
	fmt.Println(user)

	// --------------------------------------------------------
	user, err = dbDAO.GetUserByTgId(userTgId)
	if err != nil {
		log.Fatal("Failed to GET user from DB\t", err)
	}
	fmt.Println(user)

	// --------------------------------------------------------
	newUserName := "Test user"
	user.Name = &newUserName
	user, err = dbDAO.ChangeUser(user)
	if err != nil {
		log.Fatal("Failed to change user in DB\t", err)
	}
	fmt.Println(user)

	// --------------------------------------------------------
	id, err := dbDAO.DeleteUser(user.Id)
	if err != nil {
		log.Fatal("Failed to delete user from DB\t", err)
	}
	fmt.Println("User was deleted ", id)
}

func testThemes(dbDAO model.DAO) {
	userId := 13
	theme := &model.Theme{
		Title:       "Test theme 666",
		MainThemeId: nil,
		UserId:      &userId,
	}
	// --------------------------------------------------------
	theme, err := dbDAO.CreateTheme(*theme)
	if err != nil {
		log.Fatal("Failed to create theme\t", err)
	}
	log.Println(theme)

	// --------------------------------------------------------
	var themes []*model.Theme
	themes, err = dbDAO.GetAllUserThemes(userId)
	log.Println(themes)

	// --------------------------------------------------------
	theme.Title = "Test theme 3"
	newMainThemeId := 4
	theme.MainThemeId = &newMainThemeId
	theme, err = dbDAO.ChangeTheme(theme)
	log.Println("Theme successfully changed ", theme)

	// --------------------------------------------------------
	themeId, err := dbDAO.DeleteTheme(theme.Id)
	log.Println("Successfully deleted theme ", themeId)
}

func testNotes(dbDAO model.DAO) {
	userId := 13
	themeId := 1
	newNote := &model.Note{
		Title: "Test Note 2",
		Text: "Test Note 2 text text text text text text text text" +
			"\ntexttexttexttext" +
			"\ntexttexttext",
		NoteTypeId: model.SimpleNote,
		ThemeId:    themeId,
		UserId:     userId,
	}
	// --------------------------------------------------------
	newNote, err := dbDAO.CreateNote(*newNote)
	if err != nil {
		log.Fatal("Failed to create note\t", err)
	}
	log.Println(newNote)

	// --------------------------------------------------------
	var allNotes []*model.Note
	allNotes, err = dbDAO.GetAllUserNotesByTheme(userId, themeId)
	log.Println("All notes:\t", allNotes)

	// --------------------------------------------------------
	newNote.Text = "Test note 2 text text text"
	newNote.NoteTypeId = model.Question
	newNote, err = dbDAO.ChangeNote(newNote)
	log.Println("Theme successfully changed\t", newNote)

	// --------------------------------------------------------
	themeId, err = dbDAO.DeleteNote(newNote.Id)
	log.Println("Successfully deleted note ", themeId)
}
