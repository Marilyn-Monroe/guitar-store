package impl

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"guitarStore/internal/model"
	"guitarStore/internal/repository"
	"guitarStore/internal/service"
)

func NewGuitarServiceImpl(guitarRepository *repository.GuitarRepository, reviewRepository *repository.ReviewRepository) service.GuitarService {
	return &guitarServiceImpl{
		guitarRepository: *guitarRepository,
		reviewRepository: *reviewRepository,
	}
}

type guitarServiceImpl struct {
	guitarRepository repository.GuitarRepository
	reviewRepository repository.ReviewRepository
}

func (g guitarServiceImpl) FindAll(ctx context.Context) ([]model.GuitarModel, error) {
	guitars := make([]model.GuitarModel, 0)

	all, err := g.guitarRepository.FindAll(ctx)
	if err != nil {
		return guitars, fmt.Errorf("error finding all guitars: %w", err)
	}

	for _, guitar := range all {
		averageRating, err := g.reviewRepository.GetAverageRatingByGuitarId(ctx, guitar.Id)
		if err != nil {
			return guitars, fmt.Errorf("error getting average rating: %w", err)
		}

		guitars = append(guitars, model.GuitarModel{
			AverageRating:     &averageRating,
			CreatedAt:         &guitar.CreatedAt,
			Description:       guitar.Description,
			Id:                &guitar.Id,
			Image:             guitar.Image,
			Name:              &guitar.Name,
			Price:             &guitar.Price,
			QuantityAvailable: &guitar.QuantityAvailable,
			Sku:               &guitar.Sku,
			Strings:           &guitar.Strings,
			Type:              &guitar.Type,
		})
	}

	return guitars, nil
}

func (g guitarServiceImpl) FindById(ctx context.Context, id uuid.UUID) (guitar model.GuitarModel, err error) {
	guitarEntity, err := g.FindById(ctx, id)
	if err != nil {
		return guitar, fmt.Errorf("error finding guitar by id: %w", err)
	}

	averageRating, err := g.reviewRepository.GetAverageRatingByGuitarId(ctx, id)
	if err != nil {
		return guitar, fmt.Errorf("error getting average rating: %w", err)
	}

	guitar.AverageRating = &averageRating
	guitar.CreatedAt = guitarEntity.CreatedAt
	guitar.Description = guitarEntity.Description
	guitar.Id = guitarEntity.Id
	guitar.Image = guitarEntity.Image
	guitar.Name = guitarEntity.Name
	guitar.Price = guitarEntity.Price
	guitar.QuantityAvailable = guitarEntity.QuantityAvailable
	guitar.Sku = guitarEntity.Sku
	guitar.Strings = guitarEntity.Strings
	guitar.Type = guitarEntity.Type

	return guitar, nil
}
