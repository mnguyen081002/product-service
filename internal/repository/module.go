package repository

import (
	"productservice/config"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/repository/gormlib"
	"productservice/internal/repository/mongo"

	"go.uber.org/fx"
)

type UnitOfWork struct {
	ProductRepository           domain.ProductRepository
	ProductAttributesRepository domain.ProductAttributesRepository
}

func NewUnitOfWorkGorm() *UnitOfWork {
	return &UnitOfWork{
		ProductRepository:           gormlib.NewProductRepository(),
		ProductAttributesRepository: gormlib.NewProductAttributesRepository(),
	}
}

func NewUnitOfWorkMongo() *UnitOfWork {
	return &UnitOfWork{}
}

func NewUnitOfWork(config *config.Config) *UnitOfWork {
	if config.Database.Driver == "mongo" {
		return NewUnitOfWorkMongo()
	} else {
		return NewUnitOfWorkGorm()
	}
}

func NewRepository(config *config.Config, database infrastructure.Database) infrastructure.DatabaseTransaction {
	if config.Database.Driver == "mongo" {
		return mongo.NewMongoTransaction(database)
	} else {
		return gormlib.NewGormTransaction(database)
	}
}

var Module = fx.Options(
	fx.Provide(NewUnitOfWork),
	fx.Provide(NewRepository),
)
