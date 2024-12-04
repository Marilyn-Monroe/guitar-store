package repository

import (
	"context"
	"github.com/google/uuid"
	"guitarStore/internal/entity"
)

type GuitarRepository interface {
	FindAll(ctx context.Context) ([]entity.Guitar, error)
	FindById(ctx context.Context, id uuid.UUID) (entity.Guitar, error)
}
