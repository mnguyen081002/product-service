package domain

import (
	"context"
	"productservice/internal/api/request"
	"productservice/internal/infrastructure"
	"productservice/internal/models"
)

type TierVariationRepository interface {
	Create(db *infrastructure.Database, ctx context.Context, tierVar *models.TierVariations) (res *models.TierVariations, err error)
	Update(db *infrastructure.Database, ctx context.Context, tierVar map[string]interface{}, id string) (err error)
	GetById(db *infrastructure.Database, ctx context.Context, id string) (res *models.TierVariations, err error)
}

type TierVariationService interface {
	CreateTierVariation(ctx context.Context, req request.TierVariationCreate) (tierVar *models.TierVariations, err error)
	UpdateTierVariationOptions(ctx context.Context, req request.TierVariationUpdate, id string) (err error)
}
