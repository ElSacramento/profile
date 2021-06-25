package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	"github.com/profile/configuration"
	middleware2 "github.com/profile/middleware"
	"github.com/profile/notifier"
	"github.com/profile/notifier/grpc"
	"github.com/profile/service/api"
	"github.com/profile/storage"
	"github.com/profile/storage/postgres"
)

type Service struct {
	Cfg         configuration.Cfg
	Storage     storage.Storage
	HttpHandler *api.ServerWrapper
	EchoServer  *echo.Echo
	Notifier    notifier.Notifier

	Logger *logrus.Entry
}

func New(ctx context.Context, cfg configuration.Cfg) (*Service, error) {
	ctx, logger := middleware2.LoggerFromContext(ctx)

	store, err := postgres.New(ctx, cfg.DB, cfg.Migrations)
	if err != nil {
		return &Service{}, err
	}

	pushNotificator := grpc.New(ctx, cfg.Notify)

	e := echo.New()
	e.Use(
		middleware.Logger(), // todo: change to service logger
		middleware.Recover(),
	)

	server := api.NewServerWrapper(ctx, cfg.API, store, pushNotificator)
	server.RegisterHandlers(e)

	service := &Service{
		Cfg:         cfg,
		Storage:     store,
		Logger:      logger,
		EchoServer:  e,
		HttpHandler: server,
		Notifier:    pushNotificator,
	}
	return service, nil
}

func (s *Service) Start() error {
	s.Logger.Info("Starting service")

	if s.EchoServer == nil || s.HttpHandler == nil {
		return errors.New("impossible state: empty http server")
	}

	if s.Storage == nil {
		return errors.New("impossible state: no database")
	}

	if s.Notifier == nil {
		return errors.New("impossible state: nil notifier")
	}

	go func() {
		defer func() {
			s.Logger.Info("HTTP server goroutine stopped")
		}()

		s.Logger.Infof("Starting http on: %s", s.Cfg.API.Listen)
		err := s.EchoServer.Start(s.Cfg.API.Listen)
		if err != nil && err != http.ErrServerClosed {
			s.Logger.WithError(err).Fatalln("Failed to serve")
		}
	}()

	return nil
}

func (s *Service) Stop(ctx context.Context) error {
	s.Logger.Info("Stopping service")

	if s.EchoServer != nil {
		s.Logger.Infof("Stopping http on: %s", s.Cfg.API.Listen)
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

	if s.Notifier != nil {
		s.Notifier.Stop()
	}

	s.Logger.Info("Service stopped")
	return nil
}

func (s *Service) ForceStop() error {
	s.Logger.Info("Force stop service")

	if s.EchoServer != nil {
		err := s.EchoServer.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
