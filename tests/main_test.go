package tests

import (
	"os"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/ory/dockertest/v3"
	"github.com/sirupsen/logrus"

	"github.com/profile/tests/testutils"
)

func TestMain(m *testing.M) {
	// great tool: https://github.com/ory/dockertest/blob/v3/examples/PostgreSQL.md

	logger := logrus.New()

	pool, err := dockertest.NewPool("")
	if err != nil {
		logger.Fatalf("Could not connect to docker: %s", err)
	}

	opts := &dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13",
		Env: []string{
			"POSTGRES_USER=" + testutils.PostgresOptions.User,
			"POSTGRES_PASSWORD=" + testutils.PostgresOptions.Password,
			"POSTGRES_DB=" + testutils.PostgresOptions.Database,
		},
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(opts)
	if err != nil {
		logger.Fatalf("Could not start resource: %s", err.Error())
	}

	cleaner := func() {
		// When you're done, kill and remove the container
		if err = pool.Purge(resource); err != nil {
			logger.Fatalf("Could not purge resource: %s", err)
		}
	}
	defer cleaner()

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	var db *pg.DB
	defer func() {
		if err := db.Close(); err != nil {
			logger.WithError(err).Warn("failed to close connection to postgres")
		}
	}()

	if err = pool.Retry(func() error {
		db = pg.Connect(&testutils.PostgresOptions)
		_, err := db.Exec("select 1")
		return err
	}); err != nil {
		logger.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()
	os.Exit(code)
}
