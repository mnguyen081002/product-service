package models

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type ProductAttributes struct {
	BaseModel  `bson:",inline"`
	ProductID  uuid.UUID `json:"product_id" gorm:"column:product_id;type:uuid;not null"`
	Atrributes Attribute `json:"attributes" gorm:"column:attributes;type:jsonb"`
}

type Attribute struct {
	Name   string `json:"name"`
	Option string `json:"option"`
}

type ArrayAttribute []Attribute

// Value Marshal
func (a *ArrayAttribute) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *ArrayAttribute) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
