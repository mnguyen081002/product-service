package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	MissingContext = "x-user-id is not found in context, please add x-user-id or use .WithContext(ctx)"
)

type BaseModel struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" bson:"_id" json:"id"`
	UpdaterID *uuid.UUID `gorm:"column:updater_id;type:uuid;" bson:"updater_id" json:"-"`
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp;default:now();not null" bson:"created_at" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:timestamp;default:now();not null" bson:"updated_at" json:"-"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamp" bson:"deleted_at" json:"-"`
}
