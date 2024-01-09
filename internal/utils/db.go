package utils

import (
	"errors"
	"productservice/internal/infrastructure"

	"gorm.io/gorm"
)

func ErrNoRows(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func ErrNoRowEffected(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func MustHaveDb(db *infrastructure.Database) {
	if db == nil {
		panic("Database engine is null")
	}
}
