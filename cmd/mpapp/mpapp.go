package main

import (
	"MindPalace/internal/mindPalace/configuration"
	"MindPalace/internal/mindPalace/dal"
	"MindPalace/internal/mindPalace/http"
	"MindPalace/internal/mindPalace/model"
	logrus "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"time"

	"context"
	"os"
	"os/signal"
	"syscall"
)

const (
	DefaultLogLevel = logrus.InfoLevel
)

// TODO: изменить путь на /etc/..
var PathToConfig = "/home/reserv/GolandProjects/MindPalace/internal/mindPalace/config.yaml"

func main() {
	// setup logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(DefaultLogLevel) // default log level

	// read config file and setup logger level
	config, err := configuration.ReadConfig(PathToConfig)
	if err != nil {
		logger.WithField("reason", err).Fatal("error occurred when read config file")
	}
	logger.WithField("value", config).Info("config successfully parsed")
	lvl, err := logrus.ParseLevel(config.Logger.Level)
	if err != nil {
		lvl = DefaultLogLevel
		logger.WithField("reason", err).Warning("failed to parse log level, will be used " + DefaultLogLevel.String() + " as default")
	}
	logger.SetLevel(lvl)
	logger.Debugf("set log level to %s", lvl)

	// setup DB
	var dbDAO model.IDAO
	dbDAO, err = dal.NewPostgresDB(config, logrus.NewEntry(logger).WithField("module", "db"))
	if err != nil {
		logger.WithField("reason", err).Fatal("failed to create connection to DB")
	}
	logger.Info("successfully connected to DB")

	// setup services
	ctx, ctxDone := context.WithCancel(context.Background())
	eg, egContext := errgroup.WithContext(ctx)
	httpSerer := http.NewHttpServer(config, dbDAO, logrus.NewEntry(logger).WithField("module", "http_server"))

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
			logger.WithField("reason", err).Fatal("error handled when shutdown HTTP server")
		}
		return nil
	})

	err = eg.Wait()
	if err != nil {
		logger.Infof("received error: %v", err)
	} else {
		logger.Info("finished clean")
	}
}
