package request

import (
	"fmt"
	"productservice/internal/models"

	uuid "github.com/satori/go.uuid"
)

type TierVariationCreate struct {
	ProductID  string      `json:"product_id" validate:"required"`
	Variations []Variation `json:"variations" validate:"required"`
}

type TierVariationUpdate struct {
	ID      int      `json:"id" validate:"required"`
	Options []string `json:"options" validate:"required"`
}

type TierVariationDelete struct {
	ID        int      `json:"id" validate:"required"`
	ProductID string   `json:"product_id" validate:"required"`
	Options   []string `json:"options" validate:"required"`
}

type Variation struct {
	Name    string   `json:"name" validate:"required"`
	Options []string `json:"options" validate:"required"`
}

func ToOptionsTierVar(a []Variation) []models.Variation {
	var options []models.Variation
	for i, v := range a {
		options = append(options, models.Variation{
			ID:      i,
			Name:    v.Name,
			Options: v.Options,
		})
	}
	return options
}

// Convert array []models.Variation to array []models.ArrayProductModel
func ToArrayProductModel(a []Variation, productID string) []*models.ProductModel {
	var modelsOut = make([]*models.ProductModel, 0)
	init := 0

	ProductID := uuid.FromStringOrNil(productID)

	for {
		if init+1 >= len(a) {
			break
		}

		for i, v := range a[init].Options {
			for i2, v2 := range a[init+1].Options {
				modelsOut = append(modelsOut, &models.ProductModel{
					ProductID: ProductID,
					Name:      v + "," + v2,
					TierIndex: fmt.Sprintf("%d%d", i, i2),
				})
			}
		}
		init++
	}

	return modelsOut
}

// Convert array []models.Variation to array []models.ArrayProductModel
func ToArrayProductModelUpdate(a []Variation, productID string, lenArr int) []*models.ProductModel {
	var modelsOut = make([]*models.ProductModel, 0)
	init := 0

	ProductID := uuid.FromStringOrNil(productID)

	for {
		if init+1 >= len(a) {
			break
		}

		for _, v := range a[init].Options {
			for i2, v2 := range a[init+1].Options {
				modelsOut = append(modelsOut, &models.ProductModel{
					ProductID: ProductID,
					Name:      v + "," + v2,
					TierIndex: fmt.Sprintf("%d%d", lenArr+1, i2),
				})
			}
		}
		init++
	}

	return modelsOut
}
