package repository

import (
	"context"
	"github.com/google/uuid"
	"guitarStore/internal/entity"
)

type ReviewRepository interface {
	FindByGuitarId(ctx context.Context, guitarId uuid.UUID) ([]entity.Review, error)
	Insert(ctx context.Context, reviewEntity entity.Review) (entity.Review, error)
	GetAverageRatingByGuitarId(ctx context.Context, guitarId uuid.UUID) (float32, error)
}
