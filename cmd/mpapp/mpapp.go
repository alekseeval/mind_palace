package main

import (
	"MindPalace/internal/mindPalace/configuration"
	"MindPalace/internal/mindPalace/dal"
	"MindPalace/internal/mindPalace/model"
	"fmt"
	"github.com/jmoiron/sqlx"
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

	var userTgId int64 = 123123
	var user *model.User

	// --------------------------------------------------------
	var userName *string = nil
	user, err = dbDAO.SaveUser(model.User{
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
	newUserName := "Alekseev Andrey"
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

func initPostgresDb() (*sqlx.DB, error) {
	return nil, nil
}
