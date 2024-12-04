package entity

import (
	"github.com/google/uuid"
	"time"
)

// Review represents the review table.
type Review struct {
	Id            uuid.UUID `db:"id"`
	Advantages    *string   `db:"advantages"`
	Disadvantages *string   `db:"disadvantages"`
	Comments      *string   `db:"comments"`
	Rating        int64     `db:"rating"`
	GuitarId      uuid.UUID `db:"guitar_id"`
	CreatedAt     time.Time `db:"created_at"`
	CreatedBy     uuid.UUID `db:"created_by"`
}
