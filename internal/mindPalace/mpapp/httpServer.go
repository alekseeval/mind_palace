package mpapp

import (
	"MindPalace/internal/mindPalace/configuration"
	"MindPalace/internal/mindPalace/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
	"net/http"
)

type HttpServer struct {
	storage *model.IDAO
	echo    *echo.Echo
}

func NewHttpServer(config *configuration.Config, storage *model.IDAO) *HttpServer {
	return &HttpServer{
		storage: storage,
	}
}

func (s *HttpServer) ListenAndServe() error {
	e := echo.New()
	s.echo = e
	e.HideBanner = true
	e.GET("/authorize", s.authorize)
	err := e.Start(":1234")
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *HttpServer) ShoutDown(ctx context.Context) error {
	err := s.echo.Shutdown(ctx)
	return err
}
