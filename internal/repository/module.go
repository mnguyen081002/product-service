package repository

import (
	"productservice/config"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"

	"go.uber.org/fx"
)

type UnitOfWork struct {
	CmsProductRepository    domain.CmsProductRepository
	ClientProductRepository domain.ClientProductRepository
}

func NewUnitOfWorkMongo() *UnitOfWork {
	return &UnitOfWork{}
}

func NewUnitOfWork(config *config.Config) *UnitOfWork {
	if config.Database.Driver == "mongo" {
		return NewUnitOfWorkMongo()
	}
	return NewUnitOfWorkMongo()
}

func NewRepository(config *config.Config, database infrastructure.Database) infrastructure.DatabaseTransaction {
	return nil
}

var Module = fx.Options(
	fx.Provide(NewUnitOfWork),
	fx.Provide(NewRepository),
)
