package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"

	"github.com/profile/configuration"
	"github.com/profile/service/api"
	"github.com/profile/storage"
	"github.com/profile/storage/postgres"
)

type Service struct {
	Cfg         configuration.Cfg
	Storage     storage.Storage
	HttpHandler *api.ServerWrapper
	EchoServer  *echo.Echo

	Logger *logrus.Entry
}

func New(cfg configuration.Cfg) (*Service, error) {
	logger := logrus.New().WithField("layer", "service")

	store, err := postgres.New(cfg.DB, cfg.Migrations)
	if err != nil {
		return &Service{}, err
	}

	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	server := api.NewServerWrapper(cfg.API, store)
	server.RegisterHandlers(e)

	service := &Service{
		Cfg:         cfg,
		Storage:     store,
		Logger:      logger,
		EchoServer:  e,
		HttpHandler: server,
	}
	return service, nil
}

func (s *Service) Start() error {
	s.Logger.Infof("Starting service\n")

	if s.EchoServer == nil || s.HttpHandler == nil {
		return errors.New("impossible state: empty http server")
	}

	if s.Storage == nil {
		return errors.New("impossible state: no database")
	}

	go func() {
		defer func() {
			s.Logger.Infof("HTTP server goroutine stopped")
		}()

		s.Logger.Infof("Starting http on: %s", s.Cfg.API.Listen)
		err := s.EchoServer.Start(s.Cfg.API.Listen)
		if err != nil && err != http.ErrServerClosed {
			s.Logger.Fatalf("Failed to serve: ", err.Error())
		}
	}()

	return nil
}

func (s *Service) Stop(ctx context.Context) error {
	s.Logger.Info("Stopping service\n")

	if s.EchoServer != nil {
		log.Infof("Stopping http on: %s", s.Cfg.API.Listen)
		err := s.EchoServer.Shutdown(ctx)
		if err != nil {
			return err
		}
	}

	if s.HttpHandler != nil {
		err := s.HttpHandler.Stop(ctx)
		if err != nil {
			return nil
		}
	}

	if s.Storage != nil {
		err := s.Storage.Stop(ctx)
		if err != nil {
			return err
		}
	}

	s.Logger.Info("Service stopped\n")
	return nil
}

func (s *Service) ForceStop() error {
	s.Logger.Info("Force stop service\n")

	if s.EchoServer != nil {
		err := s.EchoServer.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
