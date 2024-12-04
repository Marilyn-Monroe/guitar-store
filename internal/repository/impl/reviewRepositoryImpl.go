package impl

import (
	"context"
	"fmt"
	"github.com/go-redis/cache/v9"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"guitarStore/internal/configuration/database"
	"guitarStore/internal/entity"
	"guitarStore/internal/repository"
	"time"
)

func NewReviewRepositoryImpl(databaseCluster database.DatabaseCluster, redisCluster *redis.ClusterClient) repository.ReviewRepository {
	return &reviewRepositoryImpl{
		databaseCluster: databaseCluster,
		cache: cache.New(&cache.Options{
			Redis:      redisCluster,
			LocalCache: cache.NewTinyLFU(1000, time.Minute),
		}),
	}
}

type reviewRepositoryImpl struct {
	databaseCluster database.DatabaseCluster
	cache           *cache.Cache
}

func (r reviewRepositoryImpl) FindByGuitarId(ctx context.Context, guitarId uuid.UUID) ([]entity.Review, error) {
	reviews := make([]entity.Review, 0)

	rows, err := r.databaseCluster.Slave().QueryContext(ctx, "SELECT id, advantages, disadvantages, comments, rating, guitar_id, created_at, created_by FROM review WHERE guitar_id = $1", guitarId)
	if err != nil {
		return reviews, fmt.Errorf("error selecting reviews: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var review entity.Review
		if err = rows.Scan(&review.Id, &review.Advantages, &review.Disadvantages, &review.Comments, &review.Rating, &review.GuitarId, &review.CreatedAt, &review.CreatedBy); err != nil {
			return reviews, fmt.Errorf("error scanning row: %w", err)
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func (r reviewRepositoryImpl) Insert(ctx context.Context, reviewEntity entity.Review) (review entity.Review, err error) {
	rows, err := r.databaseCluster.Master().QueryContext(ctx, "INSERT INTO review(id, advantages, disadvantages, comments, rating, guitar_id, created_at, created_by) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, advantages, disadvantages, comments, rating, guitar_id, created_at, created_by", reviewEntity.Id, reviewEntity.Advantages, reviewEntity.Disadvantages, reviewEntity.Comments, reviewEntity.Rating, reviewEntity.GuitarId, reviewEntity.CreatedAt, reviewEntity.CreatedBy)
	if err != nil {
		return review, fmt.Errorf("error inserting review: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&review.Id, &review.Advantages, &review.Disadvantages, &review.Comments, &review.Rating, &review.GuitarId, &review.CreatedAt, &review.CreatedBy); err != nil {
			return review, fmt.Errorf("error scanning row: %w", err)
		}
	}

	return review, nil
}

func (r reviewRepositoryImpl) GetAverageRatingByGuitarId(ctx context.Context, guitarId uuid.UUID) (averageRating float32, err error) {
	if err = r.cache.Get(ctx, guitarId.String(), &averageRating); err == nil {
		return averageRating, nil
	}

	rows, err := r.databaseCluster.Slave().QueryContext(ctx, "SELECT COALESCE(AVG(rating), 0) AS average_rating FROM review WHERE guitar_id = $1", guitarId)
	if err != nil {
		return 0, fmt.Errorf("error calculating average rating: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err = rows.Scan(&averageRating); err != nil {
			return averageRating, fmt.Errorf("error scanning average rating: %w", err)
		}
	}

	_ = r.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   guitarId.String(),
		Value: averageRating,
		TTL:   time.Hour,
	})

	return averageRating, nil
}
