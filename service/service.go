package service

import (
	"github.com/sirupsen/logrus"

	"github.com/profile/configuration"
	"github.com/profile/storage"
	"github.com/profile/storage/postgres"
)

type Service struct {
	cfg     configuration.Cfg
	storage storage.Storage

	logger *logrus.Entry
}

func New(cfg configuration.Cfg, st storage.Storage) (*Service, error) {
	logger := logrus.New().WithField("layer", "service")

	store, err := postgres.New(cfg.DB)
	if err != nil {
		return &Service{}, err
	}

	// todo server

	service := &Service{
		cfg:     cfg,
		storage: store,
		logger:  logger,
	}
	return service, nil
}
