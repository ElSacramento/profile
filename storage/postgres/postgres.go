package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"

	"github.com/profile/configuration"
	"github.com/profile/models"
	"github.com/profile/storage"
)

type DataStore struct {
	cfg    configuration.DB
	logger *logrus.Entry
	db     *pg.DB
}

func New(cfg configuration.DB) (*DataStore, error) {
	logger := logrus.New().WithField("layer", "postgres")

	opt, err := pg.ParseURL(cfg.URL)
	if err != nil {
		return nil, err
	}

	db := pg.Connect(opt)
	if err := pingLoop(db, logger); err != nil {
		return nil, err
	}

	// todo: migrations

	store := &DataStore{
		cfg:    cfg,
		logger: logger,
		db:     db,
	}
	return store, nil
}

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

func (s *DataStore) Create(obj models.User) (models.User, error) {
	_, err := s.db.Model(&obj).Insert()
	return obj, err
}

func (s *DataStore) Get(id uint64) (models.User, error) {
	obj := models.User{ID: id}
	err := s.db.Model(&obj).WherePK().Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return models.User{}, storage.ErrNotFound
		}
		return models.User{}, err
	}
	return obj, nil
}

func (s *DataStore) Update(update models.User) (models.User, error) {
	res, err := s.db.Model(&update).WherePK().Update()
	if err != nil {
		return models.User{}, err
	}
	if res.RowsAffected() <= 0 {
		return models.User{}, storage.ErrNotFound
	}
	return update, nil
}

func (s *DataStore) Delete(id uint64) (bool, error) {
	obj := models.User{ID: id}
	res, err := s.db.Model(&obj).WherePK().Delete()
	if err != nil {
		return false, err
	}
	return res.RowsAffected() > 0, nil
}
