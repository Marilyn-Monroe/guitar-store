package repository

import (
	"context"
	"guitarStore/internal/entity"
)

type PromocodeRepository interface {
	FindByCode(ctx context.Context, code string) (entity.Promocode, error)
}
