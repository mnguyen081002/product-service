package repository

import (
	"go.uber.org/fx"
	"productservice/config"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/repository/gormlib"
	"productservice/internal/repository/mongo"
)

type UnitOfWork struct {
	ProductRepository  domain.ProductRepository
	CategoryRepository domain.CategoryRepository
}

func NewUnitOfWorkGorm() *UnitOfWork {
	return &UnitOfWork{
		ProductRepository:  gormlib.NewProductRepository(),
		CategoryRepository: gormlib.NewCategoryRepository(),
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
