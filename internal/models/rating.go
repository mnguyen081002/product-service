package models

import (
	"time"

	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Rating struct {
	BaseModel    `bson:",inline"`
	RaterID      uuid.UUID      `json:"rater_id" gorm:"column:rater_id;type:uuid;not null" bson:"rater_id"`
	ProductID    uuid.UUID      `json:"product_id" gorm:"column:product_id;type:uuid;not null" bson:"product_id"`
	Comment      string         `json:"comment" gorm:"column:comment;type:varchar(255);not null" bson:"comment"`
	Rating       int64          `json:"rating" gorm:"column:rating;type:bigint;not null" bson:"rating"`
	Images       pq.StringArray `json:"images" gorm:"column:images;type:varchar(255)[]" bson:"images"`
	TotalLike    int64          `json:"total_like" gorm:"column:total_like;type:bigint;not null;default:0" bson:"total_like"`
	TotalDislike int64          `json:"total_dislike" gorm:"column:total_dislike;type:bigint;not null;default:0" bson:"total_dislike"`
}

func (u *Rating) MarshalBSON() ([]byte, error) {
	if u.CreatedAt.IsZero() {
		u.ID = uuid.NewV4()
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()

	type my Rating
	return bson.Marshal((*my)(u))
}
