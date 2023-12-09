package models

import (
	"productservice/internal/api_errors"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

const (
	MissingContext = "x-user-id is not found in context, please add x-user-id or use .WithContext(ctx)"
)

type BaseModel struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" bson:"_id"`
	UpdaterID *uuid.UUID `gorm:"column:updater_id;type:uuid;" bson:"updater_id"`
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp;default:now();not null" bson:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:timestamp;default:now();not null" bson:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamp" bson:"deleted_at"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	updater_id := tx.Statement.Context.Value("x-user-id")
	if updater_id != nil {
		i, err := uuid.FromString(updater_id.(string))
		if err != nil {
			return errors.New(api_errors.ErrInvalidUserID)
		}
		b.UpdaterID = &i
		return nil
	}

	return errors.New(MissingContext)
}

func (b *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	updater_id := tx.Statement.Context.Value("x-user-id")
	if updater_id != nil {
		i, err := uuid.FromString(updater_id.(string))
		if err != nil {
			return errors.New(api_errors.ErrInvalidUserID)
		}
		b.UpdaterID = &i
		return nil
	}

	return errors.New(MissingContext)
}

func (b *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	updater_id := tx.Statement.Context.Value("x-user-id")
	if updater_id != nil {
		i, err := uuid.FromString(updater_id.(string))
		if err != nil {
			return errors.New(api_errors.ErrInvalidUserID)
		}
		b.UpdaterID = &i
		return nil
	}

	return errors.New(MissingContext)
}
