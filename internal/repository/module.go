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
	ProductAttributesRepository domain.ProductAttributesRepository
	TierVariationRepository     domain.TierVariationRepository
	ProductModelRepository      domain.ProductModelRepository
	ProductRepository           domain.ProductRepository
	CategoryRepository          domain.CategoryRepository
	RatingRepository            domain.RatingRepository
	OrderRepository             domain.OrderRepository
}

func NewUnitOfWorkGorm() *UnitOfWork {
	return &UnitOfWork{
		ProductRepository:           gormlib.NewProductRepository(),
		ProductAttributesRepository: gormlib.NewProductAttributesRepository(),
		TierVariationRepository:     gormlib.NewTierVariationRepository(),
		ProductModelRepository:      gormlib.NewProductModelRepository(),
		CategoryRepository:          gormlib.NewCategoryRepository(),
		RatingRepository:            gormlib.NewRatingRepository(),
		OrderRepository:             gormlib.NewOrderRepository(),
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
