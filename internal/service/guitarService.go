package service

import (
	"context"
	"github.com/google/uuid"
	"guitarStore/internal/model"
)

type GuitarService interface {
	FindAll(ctx context.Context) ([]model.GuitarModel, error)
	FindById(ctx context.Context, id uuid.UUID) (model.GuitarModel, error)
}
