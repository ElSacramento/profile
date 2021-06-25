package tests

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	"github.com/profile/middleware"
	"github.com/profile/tests/testutils"
)

// great tool: https://github.com/ory/dockertest/blob/v3/examples/PostgreSQL.md
func TestMain(m *testing.M) {
	_, logger := middleware.LoggerFromContext(context.Background())

	pool, err := dockertest.NewPool("")
	if err != nil {
		logger.WithError(err).Fatalln("Could not connect to docker")
	}

	opts := &dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13",
		Env: []string{
			"POSTGRES_USER=" + testutils.PostgresOptions.User,
			"POSTGRES_PASSWORD=" + testutils.PostgresOptions.Password,
			"POSTGRES_DB=postgres", // default
		},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432/tcp": {{HostIP: "0.0.0.0", HostPort: "5432"}},
		},
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(opts)
	if err != nil {
		logger.WithError(err).Fatalln("Could not start resource")
	}

	var db *pg.DB
	stopper := func() {
		if err := db.Close(); err != nil {
			logger.WithError(err).Warn("Failed to close connection to postgres")
		}
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err = pool.Retry(func() error {
		db = pg.Connect(&testutils.PostgresOptions)
		_, err := db.Exec("select 1")
		return err
	}); err != nil {
		logger.WithError(err).Fatalln("Could not connect to docker")
	}

	// todo: cleaning doesn't work during debug, need to fix it
	cleaner := func() {
		// When you're done, kill and remove the container
		logger.Info("Cleaning resource")
		if err = pool.Purge(resource); err != nil {
			logger.WithError(err).Fatalln("Could not purge resource")
		}
	}

	// Handle SIGINT and SIGTERM.
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s := <-gracefulStop
		logger.Infof("Got signal: %v", s)

		stopper()
		cleaner()
		os.Exit(1)
	}()

	code := m.Run()
	stopper()
	cleaner()
	os.Exit(code)
}
