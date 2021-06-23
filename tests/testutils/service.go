package testutils

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"github.com/profile/configuration"
	"github.com/profile/service"
	"github.com/profile/service/api"
	"github.com/profile/storage"
	"github.com/profile/storage/postgres"
)

var PostgresOptions = pg.Options{
	Addr:            "localhost:5432",
	User:            "root",
	Password:        "toor",
	Database:        "test_db",
	ApplicationName: "test",
}

var PgURL = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&application_name=%s",
	PostgresOptions.User, PostgresOptions.Password, PostgresOptions.Addr, PostgresOptions.Database, PostgresOptions.ApplicationName)

func NewTestService(cfg configuration.Cfg, store storage.Storage) (*service.Service, error) {
	logger := logrus.New().WithField("layer", "service")

	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	server := api.NewServerWrapper(cfg.API, store)
	server.RegisterHandlers(e)

	srv := &service.Service{
		Cfg:         cfg,
		Storage:     store,
		Logger:      logger,
		EchoServer:  e,
		HttpHandler: server,
	}
	return srv, nil
}

func CreateTestDatabase(cfg configuration.DB) error {
	opt, err := pg.ParseURL(cfg.URL)
	if err != nil {
		return err
	}

	db := pg.Connect(opt)
	_, err = db.Exec("CREATE DATABASE %s", opt.Database)
	if err != nil {
		return err
	}
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}

func RunWithServer(t *testing.T, fn func()) {
	ctx := context.Background()
	cfg := configuration.Cfg{
		API:         configuration.API{Listen: ":8080"},
		DB:          configuration.DB{URL: PgURL},
		StopTimeout: time.Second * 60,
		Migrations:  configuration.Migrations{Directory: "../storage/postgres/migrations"},
	}

	// err := CreateTestDatabase(cfg.DB)
	// require.NoError(t, err)

	store, err := postgres.New(cfg.DB, cfg.Migrations)
	require.NoError(t, err)

	svc, err := NewTestService(cfg, store)
	require.NoError(t, err)

	err = svc.Start()
	require.NoError(t, err)

	defer func() {
		err := svc.Stop(ctx)
		require.NoError(t, err)
	}()
	fn()
}
