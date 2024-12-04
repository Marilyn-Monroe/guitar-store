package impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"guitarStore/internal/configuration/database"
	"guitarStore/internal/entity"
	"guitarStore/internal/repository"
	"log"
)

func NewCartItemRepositoryImpl(databaseCluster database.DatabaseCluster) repository.CartItemRepository {
	return &cartItemRepositoryImpl{
		databaseCluster: databaseCluster,
	}
}

type cartItemRepositoryImpl struct {
	databaseCluster database.DatabaseCluster
}

func (c cartItemRepositoryImpl) FindByUserId(ctx context.Context, userId uuid.UUID) ([]entity.CartItem, error) {
	cartItems := make([]entity.CartItem, 0)

	rows, err := c.databaseCluster.Slave().QueryContext(ctx, "SELECT user_id, guitar_id, quantity FROM cart_item WHERE user_id = $1", userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println(err.Error())
		return cartItems, fmt.Errorf("error selecting cart items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var cartItem entity.CartItem
		if err = rows.Scan(&cartItem.UserId, &cartItem.GuitarId, &cartItem.Quantity); err != nil {
			return cartItems, fmt.Errorf("error scanning row: %w", err)
		}
		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil
}

func (c cartItemRepositoryImpl) Upsert(ctx context.Context, cartItemEntity entity.CartItem) (cartItem entity.CartItem, err error) {
	rows, err := c.databaseCluster.Master().QueryContext(ctx, "INSERT INTO cart_item(user_id, guitar_id, quantity) VALUES($1, $2, $3) ON CONFLICT (user_id, guitar_id) DO UPDATE SET quantity = EXCLUDED.quantity RETURNING user_id, guitar_id, quantity", cartItemEntity.UserId, cartItemEntity.GuitarId, cartItemEntity.Quantity)
	if err != nil {
		return cartItem, fmt.Errorf("error upserting cart item: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&cartItem.UserId, &cartItem.GuitarId, &cartItem.Quantity); err != nil {
			return cartItem, fmt.Errorf("error scanning row: %w", err)
		}
	}

	return cartItem, nil
}
