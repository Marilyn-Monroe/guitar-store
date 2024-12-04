package entity

import (
	"github.com/google/uuid"
	"time"
)

// Guitar represents the guitar table.
type Guitar struct {
	Id                uuid.UUID  `db:"id"`
	Name              string     `db:"name"`
	Description       *string    `db:"description"`
	Sku               string     `db:"sku"`
	Price             int64      `db:"price"`
	Image             *string    `db:"image"`
	Type              string     `db:"type"`
	Strings           int64      `db:"strings"`
	QuantityAvailable int64      `db:"quantity_available"`
	CreatedAt         time.Time  `db:"created_at"`
	ModifiedAt        time.Time  `db:"modified_at"`
	DeletedAt         *time.Time `db:"deleted_at"`
}
