package main

import (
	"MindPalace/internal/tgbot"
	"MindPalace/internal/tgbot/configuration"
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	bot := tgbot.NewTelegramBot(config, logger.WithField("app", "Telegram Bot"))

	go func() {
		err = bot.Run()
		if err != nil {
			logger.Fatal("Failed to start Telegram Bot")
		}
	}()

	exitChl := make(chan os.Signal, 1)
	signal.Notify(exitChl, syscall.SIGINT, syscall.SIGTERM)
	<-exitChl // Wait os signal
	ctxWithTimeOut, cf := context.WithTimeout(context.Background(), 2*time.Second)
	defer cf()
	err = bot.Shutdown(ctxWithTimeOut)
	if err != nil {
		logger.WithField("error", err).Fatal("Failed to stop bot")
	}
	logger.Info("App was gracefully shutdown")
}
