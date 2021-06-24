package tests

import (
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/sirupsen/logrus"

	"github.com/profile/tests/testutils"
)

// great tool: https://github.com/ory/dockertest/blob/v3/examples/PostgreSQL.md
func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		logrus.WithError(err).Fatalln("Could not connect to docker")
	}

	opts := &dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13",
		Env: []string{
			"POSTGRES_USER=" + testutils.PostgresOptions.User,
			"POSTGRES_PASSWORD=" + testutils.PostgresOptions.Password,
			"POSTGRES_DB=" + testutils.PostgresOptions.Database,
		},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432/tcp": {{HostIP: "0.0.0.0", HostPort: "5432"}},
		},
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(opts)
	if err != nil {
		logrus.WithError(err).Fatalln("Could not start resource")
	}

	var db *pg.DB
	defer func() {
		if err := db.Close(); err != nil {
			logrus.WithError(err).Warn("Failed to close connection to postgres")
		}
	}()

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err = pool.Retry(func() error {
		db = pg.Connect(&testutils.PostgresOptions)
		_, err := db.Exec("select 1")
		return err
	}); err != nil {
		logrus.WithError(err).Fatalln("Could not connect to docker")
	}

	cleaner := func() {
		// When you're done, kill and remove the container
		logrus.Info("Cleaning resource")
		if err = pool.Purge(resource); err != nil {
			logrus.WithError(err).Fatalln("Could not purge resource")
		}
	}

	// Handle SIGINT and SIGTERM.
	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s := <-gracefulStop
		logrus.Infof("Got signal: %v", s)

		cleaner()
		os.Exit(1)
	}()

	code := m.Run()
	cleaner()
	os.Exit(code)
}
