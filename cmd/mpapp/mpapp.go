package main

import (
	"MindPalace/internal/mindPalace/configuration"
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
	fmt.Println(config)
}

func initPostgresDb() (*sqlx.DB, error) {
	return nil, nil
}
