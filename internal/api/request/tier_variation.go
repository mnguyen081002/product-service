package request

import (
	"fmt"
	"productservice/internal/models"

	uuid "github.com/satori/go.uuid"
)

// type CreateTierVariation struct {
// 	ProductID  string      `json:"product_id" validate:"required"`
// 	Variations []Variation `json:"variations" validate:"required"`
// }

type UpdateTierVariation struct {
	ID      int      `json:"id" validate:"required"`
	Options []string `json:"options" validate:"required"`
}

type DeleteTierVariation struct {
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
func ToArrayProductModel(variant []Variation, productID string) []*models.ProductModel {
	var modelsOut = make([]*models.ProductModel, 0)
	pid := uuid.FromStringOrNil(productID)
	var recursive func(int, models.ProductModel)
	recursive = func(i int, p models.ProductModel) {
		if i >= len(variant) {
			newp := p
			modelsOut = append(modelsOut, &newp)
			return
		}

		for index, v := range variant[i].Options {
			p.Name += "," + v
			p.TierIndex += fmt.Sprintf("%d", index)
			recursive(i+1, p)
			p.Name = p.Name[:len(p.Name)-len(v)-1]
			p.TierIndex = p.TierIndex[:len(p.TierIndex)-1]
		}
	}

	for i, v := range variant[0].Options {
		recursive(1, models.ProductModel{
			ProductID:           pid,
			Name:                v,
			TierIndex:           fmt.Sprintf("%d", i),
			Sold:                0,
			Price:               0,
			PriceBeforeDiscount: 0,
			Stock:               0,
		})
	}

	return modelsOut
}

// Convert array []models.Variation to array []models.ArrayProductModel
func ToArrayProductModelUpdate(variant []Variation, productID string, lenArr int) []*models.ProductModel {
	var modelsOut = make([]*models.ProductModel, 0)

	ProductID := uuid.FromStringOrNil(productID)

	// init := 0
	// for {
	// 	if init+1 >= len(a) {
	// 		break
	// 	}

	// 	for _, v := range a[init].Options {
	// 		for i2, v2 := range a[init+1].Options {
	// 			modelsOut = append(modelsOut, &models.ProductModel{
	// 				ProductID: ProductID,
	// 				Name:      v + "," + v2,
	// 				TierIndex: fmt.Sprintf("%d%d", lenArr+1, i2),
	// 			})
	// 		}
	// 	}
	// 	init++
	// }

	for i, v := range variant[0].Options {
		for _, v2 := range variant[1].Options {
			t := &models.ProductModel{
				ProductID: ProductID,
				Name:      v + "," + v2,
				TierIndex: fmt.Sprintf("%d%d", i, lenArr),
			}

			modelsOut = append(modelsOut, t)
		}
	}

	return modelsOut
}
