package model

type PromocodeModel struct {
	Code           *string `json:"code,omitempty"`
	Description    *string `json:"description,omitempty"`
	DiscountAmount *int64  `json:"discount_amount,omitempty"`
	ExpiresIn      *int64  `json:"expires_in,omitempty"`
	MaxUsage       *int64  `json:"max_usage,omitempty"`
	Name           *string `json:"name,omitempty"`
}
