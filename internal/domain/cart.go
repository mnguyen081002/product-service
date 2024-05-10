package domain

import (
	"context"
	"productservice/internal/models"

	uuid "github.com/satori/go.uuid"
)

type CartService interface {
	GetUserCart(ctx context.Context, userID uuid.UUID) (*models.Cart, error)
	AddItemToCart(ctx context.Context, userID uuid.UUID, item models.CartItem) (*models.Cart, error)
	UpdateItemInCart(ctx context.Context, userID uuid.UUID, item models.CartItem) (*models.Cart, error)
	RemoveItemFromCart(ctx context.Context, userID uuid.UUID, productModelID uuid.UUID) (*models.Cart, error)
}
