package mpapp

import (
	"MindPalace/internal/mindPalace/configuration"
	"MindPalace/internal/mindPalace/model"
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
	return nil
}

func (s *HttpServer) ShoutDown() error {
	return nil
}
