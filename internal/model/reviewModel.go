package model

import openapi_types "github.com/oapi-codegen/runtime/types"

type ReviewModel struct {
	Advantages    *string            `json:"advantages,omitempty"`
	Comments      *string            `json:"comments,omitempty"`
	Disadvantages *string            `json:"disadvantages,omitempty"`
	GuitarId      openapi_types.UUID `json:"guitar_id"`
	Rating        int64              `json:"rating"`
}
