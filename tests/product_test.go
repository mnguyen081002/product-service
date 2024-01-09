package main

import (
	"go.uber.org/zap"
	"productservice/internal/api/request"
	"productservice/internal/messaging/message"
	"productservice/internal/messaging/subscriber"
	"sync"
	"testing"
	"time"
)

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
	subscriber.NewUpdateProductSubscribe(cmsProductService, logger, config)
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
	r, err := cmsProductService.GetProductByID(ctx, p.ID.String())
	if err != nil {
		t.Errorf("Error get product by id %+v", err)
	}

	if r.Quantity != 0 {
		t.Error("Quantity is not 0", p.ID.String())
	}
}
