package model

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
	"time"
)

type GuitarModel struct {
	AverageRating     *float32            `json:"average_rating,omitempty"`
	CreatedAt         *time.Time          `json:"created_at,omitempty"`
	Description       *string             `json:"description,omitempty"`
	Id                *openapi_types.UUID `json:"id,omitempty"`
	Image             *string             `json:"image,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Price             *int64              `json:"price,omitempty"`
	QuantityAvailable *int64              `json:"quantity_available,omitempty"`
	Sku               *string             `json:"sku,omitempty"`
	Strings           *int64              `json:"strings,omitempty"`
	Type              *string             `json:"type,omitempty"`
}
