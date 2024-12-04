package impl

import (
	"context"
	"fmt"
	"guitarStore/internal/model"
	"guitarStore/internal/repository"
	"guitarStore/internal/service"
	"time"
)

func NewPromocodeServiceImpl(promocodeRepository *repository.PromocodeRepository) service.PromocodeService {
	return &promocodeServiceImpl{
		promocodeRepository: *promocodeRepository,
	}
}

type promocodeServiceImpl struct {
	promocodeRepository repository.PromocodeRepository
}

func (p promocodeServiceImpl) FindByCode(ctx context.Context, code string) (promocode model.PromocodeModel, err error) {
	promocodeRaw, err := p.promocodeRepository.FindByCode(ctx, code)
	if err != nil {
		return promocode, fmt.Errorf("error find promocode by code: %w", err)
	}

	promocode.Code = &promocodeRaw.Code
	promocode.Description = promocodeRaw.Description
	promocode.DiscountAmount = &promocodeRaw.DiscountAmount

	if promocodeRaw.ExpiredAt != nil {
		seconds := secondsUntil(*promocodeRaw.ExpiredAt)
		promocode.ExpiresIn = &seconds
	}

	promocode.MaxUsage = promocodeRaw.MaxUsage
	promocode.Name = &promocodeRaw.Name

	return promocode, nil
}

func secondsUntil(expiredAt time.Time) int64 {
	duration := time.Until(expiredAt)

	if duration <= 0 {
		return 0
	}

	return int64(duration.Seconds())
}
