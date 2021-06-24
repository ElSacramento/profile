package postgres

import (
	"context"

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

func New(cfg configuration.DB, migrations configuration.Migrations) (*DataStore, error) {
	logger := logrus.WithField("layer", "postgres")

	opt, err := pg.ParseURL(cfg.URL)
	if err != nil {
		return nil, err
	}

	db := pg.Connect(opt)
	if err := pingLoop(db, logger); err != nil {
		return nil, err
	}

	// todo: flag for run migrations
	if err := runMigrations(db, logger, migrations.Directory); err != nil {
		return nil, err
	}

	store := &DataStore{
		cfg:    cfg,
		logger: logger,
		db:     db,
	}
	return store, nil
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

func (s *DataStore) List(filter models.Filter) ([]models.User, error) {
	var result []models.User
	query := s.db.Model(&result).Order("id ASC")

	if filter.Country != "" {
		query = query.Where("country=?", filter.Country)
	}
	err := query.Select()
	return result, err
}

func (s *DataStore) Stop(ctx context.Context) error {
	return s.db.Close()
}
