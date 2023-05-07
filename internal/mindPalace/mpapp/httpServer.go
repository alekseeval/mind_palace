package mpapp

import (
	"MindPalace/internal/mindPalace/configuration"
	"MindPalace/internal/mindPalace/model"
	log "github.com/sirupsen/logrus"
)

type HttpServer struct {
	storage *model.IDAO
}

func NewHttpServer(config *configuration.Config, storage *model.IDAO) *HttpServer {
	return &HttpServer{
		storage: storage,
	}
}

func (s *HttpServer) ListenAndServe() error {
	log.Info("Http server was successfully started")
	return nil
}

func (s *HttpServer) ShoutDown() error {
	return nil
}
