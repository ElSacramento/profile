package testutils

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/require"

	"github.com/profile/configuration"
	middleware2 "github.com/profile/middleware"
	"github.com/profile/notifier/grpc"
	"github.com/profile/service"
	"github.com/profile/service/api"
	"github.com/profile/storage"
	"github.com/profile/storage/postgres"
)

var PostgresOptions = pg.Options{
	Addr:            "localhost:5432",
	User:            "root",
	Password:        "toor",
	Database:        "postgres", // default
	ApplicationName: "test",
}

var pgURL = func(dbName string) string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&application_name=%s",
		PostgresOptions.User, PostgresOptions.Password, PostgresOptions.Addr, dbName, PostgresOptions.ApplicationName)
}

var (
	serverAddr     = "localhost:8080"
	testDBName     = "test_db"
	subscriberAddr = ":50051"
)

func newTestService(ctx context.Context, cfg configuration.Cfg, store storage.Storage) (*service.Service, error) {
	ctx, logger := middleware2.LoggerFromContext(ctx)

	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	// todo: mock version
	pushNotificator := grpc.New(ctx, cfg.Notify)

	server := api.NewServerWrapper(ctx, cfg.API, store, pushNotificator)
	server.RegisterHandlers(e)

	srv := &service.Service{
		Cfg:         cfg,
		Storage:     store,
		Logger:      logger,
		EchoServer:  e,
		HttpHandler: server,
		Notifier:    pushNotificator,
	}
	return srv, nil
}

func createDatabase() error {
	db := pg.Connect(&PostgresOptions)
	_, err := db.Exec("create database " + testDBName)
	if err != nil {
		return err
	}
	if err := db.Close(); err != nil {
		return err
	}

	return nil
}

func dropDatabase() error {
	db := pg.Connect(&PostgresOptions)
	_, err := db.Exec("drop database " + testDBName)
	if err != nil {
		return err
	}
	if err := db.Close(); err != nil {
		return err
	}

	return nil
}

func RunWithServer(t *testing.T, fn func()) {
	ctx, _ := middleware2.LoggerFromContext(context.Background())

	err := createDatabase()
	require.NoError(t, err)

	defer func() {
		err := dropDatabase()
		require.NoError(t, err)
	}()

	cfg := configuration.Cfg{
		API:         configuration.API{Listen: serverAddr},
		DB:          configuration.DB{URL: pgURL(testDBName)},
		StopTimeout: time.Second * 60,
		Migrations:  configuration.Migrations{Directory: "../storage/postgres/migrations"},
	}

	store, err := postgres.New(ctx, cfg.DB, cfg.Migrations)
	require.NoError(t, err)

	svc, err := newTestService(ctx, cfg, store)
	require.NoError(t, err)

	err = svc.Start()
	require.NoError(t, err)

	defer func() {
		err := svc.Stop(ctx)
		require.NoError(t, err)
	}()
	fn()
}

func RunWithGRPCServer(t *testing.T, fn func()) {
	ctx, _ := middleware2.LoggerFromContext(context.Background())

	err := createDatabase()
	require.NoError(t, err)

	defer func() {
		err := dropDatabase()
		require.NoError(t, err)
	}()

	cfg := configuration.Cfg{
		API:         configuration.API{Listen: serverAddr},
		DB:          configuration.DB{URL: pgURL(testDBName)},
		StopTimeout: time.Second * 60,
		Migrations:  configuration.Migrations{Directory: "../storage/postgres/migrations"},
		Notify:      []string{subscriberAddr},
	}

	var subscriber *server
	go func() {
		subscriber = startGRPCServer(subscriberAddr)
	}()

	defer func() {
		if subscriber != nil {
			subscriber.stopGRPCServer()
		}
	}()

	store, err := postgres.New(ctx, cfg.DB, cfg.Migrations)
	require.NoError(t, err)

	svc, err := newTestService(ctx, cfg, store)
	require.NoError(t, err)

	err = svc.Start()
	require.NoError(t, err)

	defer func() {
		err := svc.Stop(ctx)
		require.NoError(t, err)
	}()

	fn()
}
