package main

import (
	"context"
	config2 "productservice/config"
	"productservice/internal/api/request"
	"productservice/internal/domain"
	"productservice/internal/infrastructure"
	"productservice/internal/lib"
	"productservice/internal/messaging/message"
	"productservice/internal/messaging/producer"
	"productservice/internal/messaging/subscriber"
	"productservice/internal/repository"
	service "productservice/internal/services"
	"go.uber.org/zap"
	"os"
	"sync"
	"testing"
	"time"
)

var (
	db                 infrastructure.Database
	cmsProductService  domain.CmsProductService
	cmsProductProducer producer.CmsProductProducer
	ctx                context.Context
	config             *config2.Config
	logger             *zap.Logger
)

func setup() {
	config = config2.NewConfig("/../config/config.yaml")()
	logger = lib.NewZapLogger(config)
	db = infrastructure.NewDatabase(config, logger)
	cmsProductService = service.NewCmsProductService(db, nil, repository.NewUnitOfWork(config), config, logger)

	kafkaProducer := infrastructure.NewKafkaProducer()
	cmsProductProducer = producer.NewCmsProductProducer(kafkaProducer)

	ctx = context.WithValue(context.Background(), "x-user-id", "ee564790-1e10-43a0-9968-78dda6496ff9")
}

func TestDecreaseProductQuantityCurrency(t *testing.T) {
	p, err := cmsProductService.CreateProduct(ctx, request.CreateProductRequest{
		Name:        "Test",
		Description: "Test",
		Price:       0,
		Quantity:    2,
		Images:      nil,
	})
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	subscriber.NewUpdateProductSubscribe(cmsProductService, logger)
	userID := ctx.Value("x-user-id").(string)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			err := cmsProductProducer.PublishDecreaseProductQuantity(ctx, message.DecreaseProductQuantity{
				ProductID: p.ID.String(),
				Quantity:  1,
				UserID:    userID,
			})
			if err != nil {
				t.Errorf("Error get product by id %+v", err)
			}
		}()
	}

	wg.Wait()
	time.Sleep(1 * time.Second) // wait for update product
	r, err := cmsProductService.GetProductById(ctx, p.ID.String())
	if err != nil {
		t.Errorf("Error get product by id %+v", err)
	}

	if r.Quantity != 0 {
		t.Error("Quantity is not 0", p.ID.String())
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
