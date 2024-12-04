package impl

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"guitarStore/internal/entity"
	"guitarStore/internal/model"
	"guitarStore/internal/repository"
	"guitarStore/internal/service"
)

func NewCartItemServiceImpl(cartItemRepository *repository.CartItemRepository) service.CartItemService {
	return &cartItemServiceImpl{
		cartItemRepository: *cartItemRepository,
	}
}

type cartItemServiceImpl struct {
	cartItemRepository repository.CartItemRepository
}

func (c cartItemServiceImpl) FindByUserId(ctx context.Context, userId uuid.UUID) ([]model.CartItemModel, error) {
	cartItems := make([]model.CartItemModel, 0)

	cartItemsEntity, err := c.cartItemRepository.FindByUserId(ctx, userId)
	if err != nil {
		return cartItems, fmt.Errorf("error finding by user id: %w", err)
	}

	for _, cartItem := range cartItemsEntity {
		cartItems = append(cartItems, model.CartItemModel{
			GuitarId: &cartItem.GuitarId,
			Quantity: &cartItem.Quantity,
			UserId:   &cartItem.UserId,
		})
	}

	return cartItems, nil
}

func (c cartItemServiceImpl) Update(ctx context.Context, cartItemModel model.CartItemModel) (cartItem model.CartItemModel, err error) {
	cartItemEntity, err := c.cartItemRepository.Upsert(ctx, entity.CartItem{
		GuitarId: *cartItemModel.GuitarId,
		Quantity: *cartItemModel.Quantity,
		UserId:   *cartItemModel.UserId,
	})
	if err != nil {
		return cartItem, fmt.Errorf("error upserting cart item: %w", err)
	}

	cartItem.GuitarId = &cartItemEntity.GuitarId
	cartItem.Quantity = &cartItemEntity.Quantity
	cartItem.UserId = &cartItemEntity.UserId

	return cartItem, nil
}
