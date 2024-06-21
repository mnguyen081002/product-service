package utils

import (
	"productservice/internal/infrastructure"
)

// func ErrNoRows(err error) bool {
// 	return errors.Is(err, pgx.Err)
// }

// func ErrNoRowEffected(err error) bool {
// 	return errors.Is(err, gorm.ErrRecordNotFound)
// }

func MustHaveDb(db *infrastructure.Database) {
	if db == nil {
		panic("Database engine is null")
	}
}
