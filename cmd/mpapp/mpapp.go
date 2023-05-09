package main

import (
	"MindPalace/internal/mindPalace/configuration"
	"MindPalace/internal/mindPalace/dal"
	"MindPalace/internal/mindPalace/model"
	"MindPalace/internal/mindPalace/mpapp"
	"context"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
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
		log.WithField("reason", err).Fatal("error occurred when read config file")
	}
	lvl, err := log.ParseLevel(config.Logger.Level)
	if err != nil {
		log.WithField("reason", err).Fatal("failed to parse log level")
	}
	log.SetLevel(lvl)
	log.Debugf("Set log level to %s", lvl)

	// setup DB
	var dbDAO model.IDAO
	dbDAO, err = dal.NewPostgresDB(config)
	if err != nil {
		log.WithField("reason", err).Fatal("failed to create connection to DB")
	}
	log.Info("successfully connected to DB")

	// setup services
	ctx, done := context.WithCancel(context.Background())
	eg, egContext := errgroup.WithContext(ctx)
	exitChl := make(chan os.Signal, 1)
	signal.Notify(exitChl, syscall.SIGINT, syscall.SIGTERM)

	eg.Go(func() error {
		httpSerer := mpapp.NewHttpServer(config, &dbDAO)
		err = httpSerer.ListenAndServe()
		if err != nil {
			return err
		}
		<-egContext.Done()
		err = httpSerer.ShoutDown()
		return err
	})

	eg.Go(func() error {
		select {
		case <-exitChl:
			done()
		case <-egContext.Done():
			return egContext.Err()
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
