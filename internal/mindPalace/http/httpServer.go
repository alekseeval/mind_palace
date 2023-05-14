package http

import (
	"MindPalace/internal/mindPalace/configuration"
	"MindPalace/internal/mindPalace/model"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

type HttpServer struct {
	storage model.IDAO
	echo    *echo.Echo

	httpConfig *configuration.HttpConfig
}

func NewHttpServer(config *configuration.Config, storage model.IDAO) *HttpServer {
	httpServer := &HttpServer{
		storage:    storage,
		httpConfig: &config.System.Http,
	}

	// setup echo
	e := echo.New()
	e.HideBanner = true
	e.Debug = false
	e.Server.ReadTimeout = time.Duration(httpServer.httpConfig.ReadTimeout) * time.Second
	e.Server.WriteTimeout = time.Duration(httpServer.httpConfig.WriteTimeout) * time.Second
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Use(logMiddleware)
	httpServer.echo = e

	// endpoints
	e.GET("/users/:id", httpServer.getUser)
	e.GET("/telegram/users/:tg_id", httpServer.getUserByTgId)
	e.DELETE("/users/:id", httpServer.deleteUser)
	e.PATCH("/users/:id", httpServer.changeUser)
	e.POST("/users", httpServer.createUser)

	return httpServer
}

func (s *HttpServer) ListenAndServe() error {
	err := s.echo.Start(fmt.Sprintf("%s:%d", s.httpConfig.Host, s.httpConfig.Port))
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *HttpServer) ShoutDown(ctx context.Context) error {
	err := s.echo.Shutdown(ctx)
	return err
}
