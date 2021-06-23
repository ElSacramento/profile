package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

func pingLoop(db *pg.DB, logger *logrus.Entry) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeout := 60 * time.Second // can be in cfg

	for {
		err := db.Ping(context.Background())
		if err == nil {
			return nil
		}

		select {
		case <-time.After(timeout):
			return errors.New(fmt.Sprintf("db ping failed after %s timeout", timeout))
		case <-ticker.C:
			logger.Warn("db ping failed, sleep and retry")
		}
	}
}

func runMigrations(db *pg.DB, logger *logrus.Entry, directory string) error {
	collection := migrations.NewCollection()
	if err := collection.DiscoverSQLMigrations(directory); err != nil {
		return err
	}

	// for go_pg_migrations
	_, _, err := collection.Run(db, "init")
	if err != nil {
		return err
	}

	oldVersion, newVersion, err := collection.Run(db, "up")
	if err != nil {
		return err
	}
	if newVersion != oldVersion {
		logger.Infof("Migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		logger.Infof("Version is %d after up migrations\n", oldVersion)
	}
	return nil
}
