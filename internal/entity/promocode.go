package entity

import (
	"github.com/google/uuid"
	"time"
)

// Promocode represents the promocode table.
type Promocode struct {
	Id             uuid.UUID  `db:"id"`
	Name           string     `db:"name"`
	Description    *string    `db:"description"`
	Code           string     `db:"code"`
	MaxUsage       *int64     `db:"max_usage"`
	DiscountAmount int64      `db:"discount_amount"`
	ExpiredAt      *time.Time `db:"expired_at"`
	CreatedAt      time.Time  `db:"created_at"`
	ModifiedAt     time.Time  `db:"modified_at"`
	DeletedAt      *time.Time `db:"deleted_at"`
}
