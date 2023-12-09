package main

import (
	"context"
	"go.uber.org/zap"
	config2 "productservice/config"
	"productservice/internal/api/request"
	"productservice/internal/infrastructure"
	"productservice/internal/lib"
	"productservice/internal/messaging/message"
	"productservice/internal/messaging/subscriber"
	"sync"
	"testing"
	"time"
)

func setup() {
	config = config2.NewConfig("/../config/config.yaml")()
	logger = lib.NewZapLogger(config)
	db = infrastructure.NewDatabase(config, logger)

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
			err := cmsProductProducer.DecreaseProductQuantity(ctx, message.DecreaseProductQuantity{
				ProductID: p.ID.String(),
				Quantity:  1,
				UserID:    userID,
			})

			logger.Debug("Decrease product quantity", zap.String("product_id", p.ID.String()), zap.Int("quantity", 1))

			if err != nil {
				t.Errorf("Error get product by id %+v", err)
			}
		}()
	}

	wg.Wait()
	time.Sleep(4 * time.Second) // wait for update product
	r, err := cmsProductService.GetProductById(ctx, p.ID.String())
	if err != nil {
		t.Errorf("Error get product by id %+v", err)
	}

	if r.Quantity != 0 {
		t.Error("Quantity is not 0", p.ID.String())
	}
}
