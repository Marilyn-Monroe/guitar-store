package entity

import "github.com/google/uuid"

// CartItem represents the cart_item table.
type CartItem struct {
	UserId   uuid.UUID `db:"user_id"`
	GuitarId uuid.UUID `db:"guitar_id"`
	Quantity int64     `db:"quantity"`
}
