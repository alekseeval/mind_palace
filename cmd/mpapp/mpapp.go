package main

import (
	"MindPalace/internal/mindPalace/configuration"
	"MindPalace/internal/mindPalace/dal"
	"MindPalace/internal/mindPalace/model"
	"MindPalace/internal/mindPalace/mpapp"
	"log"
)

var PATH_TO_CONFIG string = "/home/reserv/GolandProjects/MindPalace/internal/mindPalace/config.yaml"

func main() {
	config, err := configuration.ReadConfig(PATH_TO_CONFIG)
	if err != nil {
		log.Fatal("Error occurred when read config file\t", err)
	}

	var dbDAO model.IDAO
	dbDAO, err = dal.NewPostgresDB(config)
	if err != nil {
		log.Fatal("Failed to create connection to DB\t", err)
	}

	httpSerer := mpapp.NewHttpServer(config, &dbDAO)
	httpSerer.ListenAndServe()
	httpSerer.ShoutDown()

}
