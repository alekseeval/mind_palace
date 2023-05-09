package mpapp

import (
	"MindPalace/internal/mindPalace/configuration"
	"MindPalace/internal/mindPalace/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

type HttpServer struct {
	storage *model.IDAO
	echo    *echo.Echo

	httpConfig *configuration.HttpConfig
}

func NewHttpServer(config *configuration.Config, storage *model.IDAO) *HttpServer {
	return &HttpServer{
		storage:    storage,
		httpConfig: &config.System.Http,
	}
}

func (s *HttpServer) ListenAndServe() error {
	// setup echo
	e := echo.New()
	e.HideBanner = true
	e.Server.ReadTimeout = time.Duration(s.httpConfig.ReadTimeout)
	e.Server.WriteTimeout = time.Duration(s.httpConfig.WriteTimeout)
	s.echo = e

	// endpoints
	e.GET("/authorize", s.authorize)
	e.GET("/register", s.register)

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
