package sqlx

import "productservice/internal/infrastructure"

type SqlxRepository struct {
	infrastructure.Database
}

// WithTransaction implements infrastructure.DatabaseTransaction.
func (s SqlxRepository) WithTransaction(txFunc func(tx *infrastructure.Database) error) (err error) {
	panic("unimplemented")
}

func NewSqlxTransaction(db infrastructure.Database) infrastructure.DatabaseTransaction {
	return SqlxRepository{db}
}
