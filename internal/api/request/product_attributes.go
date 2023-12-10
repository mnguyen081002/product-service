package request

import "productservice/internal/models"

type ProductAttributesCreate struct {
	ProductID  string      `json:"product_id" validate:"required"`
	Attributes []Attribute `json:"attributes" validate:"required"`
}

type ProductAttributesUpdate struct {
	Attributes []Attribute `json:"attributes" validate:"required"`
}

type Attribute struct {
	Name   string `json:"name" validate:"required"`
	Option string `json:"option" validate:"required"`
}

func (a *ProductAttributesCreate) ConvertAttributeModel() []models.Attribute {
	return convertAttributes(a.Attributes)
}

func (a *ProductAttributesUpdate) ConvertAttributeModel() []models.Attribute {
	return convertAttributes(a.Attributes)
}

func convertAttributes(attributes []Attribute) []models.Attribute {
	var res []models.Attribute
	for _, v := range attributes {
		res = append(res, models.Attribute{
			Name:   v.Name,
			Option: v.Option,
		})
	}
	return res
}
