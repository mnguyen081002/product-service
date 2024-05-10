package models

import (
	"productservice/internal/api_errors"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Cart struct {
	Items  []CartItem `json:"items"`
	UserID uuid.UUID  `json:"user_id"`
}

func (c *Cart) AddItem(item CartItem) {
	item.Validate()
	c.Items = append(c.Items, item)
}

func (c *Cart) RemoveItem(productModelId uuid.UUID) {
	for index, item := range c.Items {
		if item.ProductModel.ID == productModelId {
			c.Items = append(c.Items[:index], c.Items[index+1:]...)
			return
		}
	}
}

func (c *Cart) UpdateItem(index int, item CartItem) {
	item.Validate()
	c.Items[index] = item
}

type CartItem struct {
	Quantity     int          `json:"quantity"`
	ProductModel ProductModel `json:"product"`
}

func (c *CartItem) Validate() error {
	if c.Quantity <= 0 {
		return errors.New(api_errors.ErrCartQuantityMustHigherThanZero)
	}
	return nil
}
