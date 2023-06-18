package main

import (
	"MindPalace/internal/tgbot/configuration"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	PathToConfig    = "/home/reserv/GolandProjects/MindPalace/internal/tgbot/config.yaml"
	DefaultLogLevel = logrus.InfoLevel
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(DefaultLogLevel)
	config, err := configuration.ReadConfig(PathToConfig)
	logger.WithField("config", *config).Info("Config was parsed")
	if err != nil {
		logger.WithField("error", err).Fatal("Failed to read configuration file")
	}
	logLevel, err := logrus.ParseLevel(config.Logger.Level)
	if err != nil {
		logger.WithField("error", err).Error("Failed to determine log level from config. Default level was set to ", DefaultLogLevel)
	} else {
		logger.SetLevel(logLevel)
		logger.Info("Log level set to ", logLevel)
	}
	fmt.Println("Hello world!")
}
