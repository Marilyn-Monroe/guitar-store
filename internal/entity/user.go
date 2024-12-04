package entity

import (
	"github.com/google/uuid"
	"time"
)

// User represents the user table.
type User struct {
	Id          uuid.UUID `db:"id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	MiddleName  string    `db:"middle_name"`
	Phone       string    `db:"phone"`
	Email       string    `db:"email"`
	CreatedAt   time.Time `db:"created_at"`
	LastLoginAt time.Time `db:"last_login_at"`
}
