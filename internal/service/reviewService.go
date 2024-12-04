package service

import (
	"context"
	"github.com/google/uuid"
	"guitarStore/internal/model"
)

type ReviewService interface {
	FindByGuitarId(ctx context.Context, guitarId uuid.UUID) ([]model.ReviewModel, error)
	Create(ctx context.Context, reviewModel model.ReviewModel) (model.ReviewModel, error)
}
