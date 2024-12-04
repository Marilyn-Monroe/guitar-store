package service

import (
	"context"
	"github.com/google/uuid"
	"guitarStore/internal/model"
)

type CartItemService interface {
	FindByUserId(ctx context.Context, userId uuid.UUID) ([]model.CartItemModel, error)
	Update(ctx context.Context, cartItemModel model.CartItemModel) (model.CartItemModel, error)
}
