package impl

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"guitarStore/internal/entity"
	"guitarStore/internal/model"
	"guitarStore/internal/repository"
	"guitarStore/internal/service"
	"time"
)

func NewReviewServiceImpl(reviewRepository *repository.ReviewRepository) service.ReviewService {
	return &reviewServiceImpl{
		reviewRepository: *reviewRepository,
	}
}

type reviewServiceImpl struct {
	reviewRepository repository.ReviewRepository
}

func (r reviewServiceImpl) FindByGuitarId(ctx context.Context, guitarId uuid.UUID) ([]model.ReviewModel, error) {
	reviews := make([]model.ReviewModel, 0)

	reviewsEntity, err := r.reviewRepository.FindByGuitarId(ctx, guitarId)
	if err != nil {
		return reviews, fmt.Errorf("error finding reviews: %w", err)
	}

	for _, reviewEntity := range reviewsEntity {
		reviews = append(reviews, model.ReviewModel{
			Advantages:    reviewEntity.Advantages,
			Comments:      reviewEntity.Comments,
			Disadvantages: reviewEntity.Disadvantages,
			GuitarId:      reviewEntity.GuitarId,
			Rating:        reviewEntity.Rating,
		})
	}

	return reviews, nil
}

func (r reviewServiceImpl) Create(ctx context.Context, reviewModel model.ReviewModel) (review model.ReviewModel, err error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return review, fmt.Errorf("error generating id: %w", err)
	}

	reviewEntity, err := r.reviewRepository.Insert(ctx, entity.Review{
		Id:            id,
		Advantages:    reviewModel.Advantages,
		Disadvantages: reviewModel.Disadvantages,
		Comments:      reviewModel.Comments,
		Rating:        reviewModel.Rating,
		GuitarId:      reviewModel.GuitarId,
		CreatedAt:     time.Now(),
		CreatedBy:     uuid.Nil,
	})
	if err != nil {
		return review, fmt.Errorf("error inserting review: %w", err)
	}

	review.Advantages = reviewEntity.Advantages
	review.Comments = reviewEntity.Comments
	review.Disadvantages = reviewEntity.Disadvantages
	review.GuitarId = reviewEntity.GuitarId
	review.Rating = reviewEntity.Rating

	return review, nil
}
