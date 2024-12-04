package service

import (
	"context"
	"guitarStore/internal/model"
)

type PromocodeService interface {
	FindByCode(ctx context.Context, code string) (model.PromocodeModel, error)
}
