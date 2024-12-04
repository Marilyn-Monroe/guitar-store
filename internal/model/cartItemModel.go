package model

import openapiTypes "github.com/oapi-codegen/runtime/types"

type CartItemModel struct {
	GuitarId *openapiTypes.UUID `json:"guitar_id,omitempty"`
	Quantity *int64             `json:"quantity,omitempty"`
	UserId   *openapiTypes.UUID `json:"user_id,omitempty"`
}
