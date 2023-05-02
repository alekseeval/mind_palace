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
		log.Fatal("Error occurred when read config file", err)
	}

	dbDAO, err := dal.NewPostgresDB(config)
	if err != nil {
		log.Fatal("Failed to create connection to DB", err)
	}
	user, err := dbDAO.SaveUser(model.User{
		Name:       "Alekseev Andrey",
		TelegramId: 123123123123,
	})
	if err != nil {
		log.Fatal("Failed to create user")
	}
	fmt.Println(user)
}

func initPostgresDb() (*sqlx.DB, error) {
	return nil, nil
}
