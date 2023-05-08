package main

import (
	"MindPalace/internal/mindPalace/configuration"
	"MindPalace/internal/mindPalace/dal"
	"MindPalace/internal/mindPalace/model"
	"MindPalace/internal/mindPalace/mpapp"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var PathToConfig = "/home/reserv/GolandProjects/MindPalace/internal/mindPalace/config.yaml"

func main() {
	// read config file and setup logger
	log.SetFormatter(&log.JSONFormatter{})
	config, err := configuration.ReadConfig(PathToConfig)
	if err != nil {
		log.WithField("reason", err).Fatal("Error occurred when read config file")
	}
	lvl, err := log.ParseLevel(config.Logger.Level)
	fmt.Println(lvl)
	if err != nil {
		log.WithField("reason", err).Fatal("Failed to parse log level")
	}

	// Setup services
	var dbDAO model.IDAO
	dbDAO, err = dal.NewPostgresDB(config)
	if err != nil {
		log.WithField("reason", err).Fatal("Failed to create connection to DB")
	}
	log.Info("Successfully connected to DB")

	exitChl := make(chan os.Signal, 1)
	signal.Notify(exitChl, syscall.SIGINT, syscall.SIGTERM)

	httpSerer := mpapp.NewHttpServer(config, &dbDAO)
	httpSerer.ListenAndServe()
	<-exitChl
	httpSerer.ShoutDown()

}
