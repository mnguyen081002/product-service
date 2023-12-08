package models

import (
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Category struct {
	BaseModel `bson:",inline"`
	Name      string    `json:"name" gorm:"column:name;type:varchar(255);not null" bson:"name"`
	ParentID  uuid.UUID `json:"parent_id" gorm:"column:parent_id;type:uuid" bson:"parent_id"`
	NoSub     bool      `json:"no_sub" gorm:"column:no_sub;type:boolean;not null;default:false" bson:"no_sub"`
}

func (u *Category) MarshalBSON() ([]byte, error) {
	if u.CreatedAt.IsZero() {
		u.ID = uuid.NewV4()
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()

	type my Category
	return bson.Marshal((*my)(u))
}
