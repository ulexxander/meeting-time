package storage

import (
	"errors"

	"gorm.io/gorm"
)

var ErrNoOrganization = errors.New("organization does not exist")

type Store struct {
	DB *gorm.DB
}

func (s *Store) mapError(err error) error {
	if err == gorm.ErrRecordNotFound {
		return ErrNoOrganization
	}
	return err
}
