package models

type TierVariations struct {
	BaseModel `bson:",inline"`
	ProductID string   `json:"product_id" gorm:"column:product_id;type:uuid;not null" bson:"product_id"`
	Name      string   `json:"name" gorm:"column:name;type:varchar(255);not null" bson:"name"`
	Options   []string `json:"options" gorm:"column:options;type:varchar(255);not null" bson:"-"`
}
