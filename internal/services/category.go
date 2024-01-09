package service

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"productservice/config"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
	"productservice/internal/repository"
	"sync"
)

type categoryService struct {
	db              infrastructure.Database
	dbTransaction   infrastructure.DatabaseTransaction
	categoryService domain.CategoryService
	ufw             *repository.UnitOfWork
	config          *config.Config
	logger          *zap.Logger
	mu              sync.Mutex
}

func NewCategoryService(
	db infrastructure.Database,
	dbTransaction infrastructure.DatabaseTransaction,
	ufw *repository.UnitOfWork,
	config *config.Config,
	logger *zap.Logger,
) domain.CategoryService {
	return &categoryService{
		db:            db,
		dbTransaction: dbTransaction,
		ufw:           ufw,
		config:        config,
		logger:        logger,
	}
}

func (c *categoryService) GetProductCategories(ctx context.Context, productID uuid.UUID) (categories []*models.Category, err error) {
	return c.ufw.CategoryRepository.GetProductCategories(&c.db, ctx, productID)
}
