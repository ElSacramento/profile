package storage

import (
	"errors"

	"github.com/profile/models"
)

type Storage interface {
	Create(obj models.User) (models.User, error)
	Get(id uint64) (models.User, error)
	Update(update models.User) (models.User, error)
	Delete(id uint64) (bool, error)
}

var ErrNotFound = errors.New("user is not found")
