package http

import (
	"MindPalace/internal/mindPalace/configuration"
	"MindPalace/internal/mindPalace/model"
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

type HttpServer struct {
	storage  model.IDAO
	echo     *echo.Echo
	logEntry *log.Entry

	httpConfig *configuration.HttpConfig
}

func NewHttpServer(config *configuration.Config, storage model.IDAO, logEntry *log.Entry) *HttpServer {
	logEntry = logEntry.WithField("app", "HTTPServer")
	httpServer := &HttpServer{
		storage:    storage,
		httpConfig: &config.System.Http,
		logEntry:   logEntry,
	}

	// setup echo
	e := echo.New()
	e.HideBanner = true
	e.Debug = false
	e.Server.ReadTimeout = time.Duration(httpServer.httpConfig.ReadTimeout) * time.Second
	e.Server.WriteTimeout = time.Duration(httpServer.httpConfig.WriteTimeout) * time.Second
	e.HTTPErrorHandler = httpServer.customHTTPErrorHandler
	e.Use(httpServer.logMiddleware)
	httpServer.echo = e

	apiV1 := e.Group("/api/v1")

	// endpoints API v1
	apiV1.GET("/users/:id", httpServer.getUser)
	apiV1.GET("/telegram/users/:tg_id", httpServer.getUserByTgId)
	apiV1.DELETE("/users/:id", httpServer.deleteUser)
	apiV1.PATCH("/users/:id", httpServer.editUser)
	apiV1.POST("/users", httpServer.createUser)
	apiV1.GET("/users", httpServer.getAllUsers)

	apiV1.POST("/themes", httpServer.createTheme)
	apiV1.GET("/themes", httpServer.getUserThemes)
	apiV1.DELETE("/themes/:id", httpServer.deleteTheme)
	apiV1.PATCH("/themes/:id", httpServer.editTheme)

	apiV1.POST("/themes/:theme_id/notes", httpServer.createNote)
	apiV1.GET("/themes/:theme_id/notes", httpServer.getNotes)
	apiV1.DELETE("/notes/:note_id", httpServer.deleteNote)
	apiV1.PATCH("/notes/:note_id", httpServer.editNote)

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
