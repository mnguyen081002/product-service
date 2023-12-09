package main

import (
	"context"
	"go.uber.org/zap"
	config2 "productservice/config"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/lib"
	"productservice/internal/messaging/producer"
	"productservice/internal/repository"
	"productservice/internal/repository/gormlib"
	service "productservice/internal/services"
)

var (
	db                 infrastructure.Database
	dbTransaction      infrastructure.DatabaseTransaction
	cmsCategoryService domain.CmsCategoryService
	ctx                context.Context
	config             *config2.Config
	logger             *zap.Logger
	cmsProductService  domain.CmsProductService
	cmsProductProducer producer.CmsProductProducer
	ufw                *repository.UnitOfWork
)

func SetUp() {
	config = config2.NewConfig("/../config/config-test.yaml")()
	logger = lib.NewZapLogger(config)
	db = infrastructure.NewDatabase(config, logger)
	dbTransaction = gormlib.NewGormTransaction(db)
	ufw = repository.NewUnitOfWork(config)
	cmsCategoryService = service.NewCmsCategoryService(db, dbTransaction, ufw, config, logger)
	cmsProductService = service.NewCmsProductService(db, nil, repository.NewUnitOfWork(config), config, logger)

	kafkaProducer := infrastructure.NewKafkaProducer()
	cmsProductProducer = producer.NewCmsProductProducer(kafkaProducer, logger)
	ctx = context.WithValue(context.Background(), "x-user-id", "ee564790-1e10-43a0-9968-78dda6496ff9")
}
