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
	Variations []Variation `json:"variations" validate:"required"`
}

// Convert to model tier variation
func (a *TierVariationCreate) ToModelTierVar() *models.TierVariations {
	var options []models.Variation
	for i, v := range a.Variations {
		options = append(options, models.Variation{
			ID:      i,
			Name:    v.Name,
			Options: v.Options,
		})
	}
	return &models.TierVariations{
		ProductID: uuid.FromStringOrNil(a.ProductID),
		Options:   options,
	}
}

type Variation struct {
	Name    string   `json:"name" validate:"required"`
	Options []string `json:"options" validate:"required"`
}

// Convert to model tier variation update
func (a *TierVariationCreate) ToModelTierVarUpdate() *models.TierVariations {
	var options []models.Variation
	for _, v := range a.Variations {
		options = append(options, models.Variation{
			Name:    v.Name,
			Options: v.Options,
		})
	}
	return &models.TierVariations{
		Options: options,
	}
}

// Convert to model tier variation update
func (a *TierVariationUpdate) ToModelTierVarUpdate() *models.TierVariations {
	var options []models.Variation
	for _, v := range a.Variations {
		options = append(options, models.Variation{
			Name:    v.Name,
			Options: v.Options,
		})
	}
	return &models.TierVariations{
		Options: options,
	}
}

func (a *TierVariationUpdate) ToOptionsTierVar() []models.Variation {
	var options []models.Variation
	for _, v := range a.Variations {
		options = append(options, models.Variation{
			Name:    v.Name,
			Options: v.Options,
		})
	}
	return options
}

// Convert array variation to array product model
func (a *TierVariationCreate) ToArrayProductModel() models.ArrayProductModel {
	var modelsOut = make(models.ArrayProductModel, 0)
	init := 0

	for {
		if init+1 >= len(a.Variations) {
			break
		}

		for i, v := range a.Variations[init].Options {
			for i2, v2 := range a.Variations[init+1].Options {
				modelsOut = append(modelsOut, models.ProductModel{
					Name:      v + "," + v2,
					TierIndex: fmt.Sprintf("%d%d", i, i2),
				})
			}
		}
		init++
	}

	return modelsOut
}
