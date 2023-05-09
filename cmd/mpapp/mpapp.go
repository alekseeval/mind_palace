package main

import (
	"MindPalace/internal/mindPalace/configuration"
	"MindPalace/internal/mindPalace/dal"
	"MindPalace/internal/mindPalace/model"
	"MindPalace/internal/mindPalace/mpapp"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"time"

	"context"
	"os"
	"os/signal"
	"syscall"
)

// TODO: изменить путь на /etc/..
var PathToConfig = "/home/reserv/GolandProjects/MindPalace/internal/mindPalace/config.yaml"

func main() {
	// setup logger
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel) // default log level

	// read config file and setup logger level
	config, err := configuration.ReadConfig(PathToConfig)
	if err != nil {
		log.WithField("reason", err).Fatal("error occurred when read config file")
	}
	lvl, err := log.ParseLevel(config.Logger.Level)
	if err != nil {
		lvl = log.InfoLevel
		log.WithField("reason", err).Error("failed to parse log level")
	}
	log.SetLevel(lvl)
	log.WithField("value", config).Debug("config successfully parsed")
	log.Debugf("set log level to %s", lvl)

	// setup DB
	var dbDAO model.IDAO
	dbDAO, err = dal.NewPostgresDB(config)
	if err != nil {
		log.WithField("reason", err).Fatal("failed to create connection to DB")
	}
	log.Info("successfully connected to DB")

	// setup services
	ctx, ctxDone := context.WithCancel(context.Background())
	eg, egContext := errgroup.WithContext(ctx)
	httpSerer := mpapp.NewHttpServer(config, dbDAO)

	eg.Go(func() error {
		err = httpSerer.ListenAndServe()
		return err
	})

	eg.Go(func() error {
		exitChl := make(chan os.Signal, 1)
		signal.Notify(exitChl, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-exitChl:
			// case when captured os signal
			ctxDone()
		case <-egContext.Done():
			// case when captured error in errgroup
			return egContext.Err()
		}
		ctxWithTimeout, ctxWithTimeoutCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer ctxWithTimeoutCancel()
		err = httpSerer.ShoutDown(ctxWithTimeout)
		if err != nil {
			log.WithField("reason", err).Fatal("error handled when shutdown HTTP server")
		}
		return nil
	})

	err = eg.Wait()
	if err != nil {
		log.Infof("received error: %v", err)
	} else {
		log.Info("finished clean")
	}
}
