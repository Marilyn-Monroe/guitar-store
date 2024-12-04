package repository

import (
	"context"
	"github.com/google/uuid"
	"guitarStore/internal/entity"
)

type CartItemRepository interface {
	FindByUserId(ctx context.Context, userId uuid.UUID) ([]entity.CartItem, error)
	Upsert(ctx context.Context, cartItemEntity entity.CartItem) (entity.CartItem, error)
}
