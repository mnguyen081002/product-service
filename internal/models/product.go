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
