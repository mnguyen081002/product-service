package models

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Variation struct {
	ID      int      `json:"id" gorm:"column:id;type:bigint" bson:"-"`
	Name    string   `json:"name"`
	Options []string `json:"options"`
}

type TierVariations struct {
	BaseModel `bson:",inline"`
	ProductID uuid.UUID      `json:"product_id" gorm:"column:product_id;type:uuid;not null" bson:"product_id"`
	Options   ArrayVariation `json:"options" gorm:"column:options;type:jsonb;not null" bson:"-"`
}

// get length of array options based on id
func GetLengthOptions(a []Variation, id int) int {
	for _, v := range a {
		if v.ID == id {
			return len(v.Options)
		}
	}

	return 0
}

type ArrayVariation []Variation

// Value Marshal
func (a ArrayVariation) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *ArrayVariation) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
