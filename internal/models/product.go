package models

import (
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Product struct {
	BaseModel   `bson:",inline"`
	Name        string         `json:"name" gorm:"column:name;type:varchar(255);not null" bson:"name"`
	Description string         `json:"description" gorm:"column:description;type:varchar(255);not null" bson:"description"`
	Price       int64          `json:"price" gorm:"column:price;type:bigint;not null" bson:"price"`
	Quantity    int64          `json:"quantity" gorm:"column:quantity;type:bigint;not null" bson:"quantity"`
	Images      pq.StringArray `json:"images" gorm:"column:images;type:varchar(255)[]" bson:"images"`
	OriPrice    int64          `json:"ori_price" gorm:"column:ori_price;type:bigint;not null;default:0" bson:"ori_price"`
	TotalSold   int64          `json:"total_sold" gorm:"column:total_sold;type:bigint;not null;default:0" bson:"total_sold"`
	MedRating   float64        `json:"med_rating" gorm:"column:med_rating;type:float;not null;default:0" bson:"med_rating"`
	RatingCount int64          `json:"rating_count" gorm:"column:rating_count;type:bigint;not null;default:0" bson:"rating_count"`
	CityID      string         `json:"city_id" gorm:"column:city_id;type:varchar(255)" bson:"city_id"`
	Category    *Category      `json:"category,omitempty"`
	CategoryID  *uuid.UUID     `json:"-" gorm:"column:category_id;type:uuid;not null" bson:"category_id"`
	Rating      []*Rating      `json:"rating,omitempty" gorm:"foreignKey:ProductID;references:ID" bson:"rating"`
}

func (u *Product) MarshalBSON() ([]byte, error) {
	if u.CreatedAt.IsZero() {
		u.ID = uuid.NewV4()
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()

	type my Product
	return bson.Marshal((*my)(u))
}
