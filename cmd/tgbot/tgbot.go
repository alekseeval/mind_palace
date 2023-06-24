package main

import (
	"MindPalace/internal/tgbot"
	"MindPalace/internal/tgbot/configuration"
	"context"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
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
	ctx, ctxDoneFunc := context.WithCancel(context.Background())
	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		err = bot.Run()
		if err != nil {
			logger.Error("Failed to start Telegram Bot")
			return err
		}
		return nil
	})

	eg.Go(func() error {
		exitChl := make(chan os.Signal, 1)
		signal.Notify(exitChl, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-exitChl:
			// case when captured os signal
			ctxDoneFunc()
		case <-egCtx.Done():
			// case when captured error in errgroup
			err = egCtx.Err()
			if err != nil {
				return err
			}
			return nil
		}
		ctxWithTimeOut, cf := context.WithTimeout(context.Background(), 2*time.Second)
		defer cf()
		err = bot.Shutdown(ctxWithTimeOut)
		if err != nil {
			logger.WithField("error", err).Fatal("Failed to stop bot")
		}
		return nil
	})

	err = eg.Wait()
	if err != nil {
		logger.WithField("err", err).Error("Error occurred")
	}

	logger.Info("Shutdown the app")
}
