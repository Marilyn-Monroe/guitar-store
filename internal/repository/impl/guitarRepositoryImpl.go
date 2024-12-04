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

func NewGuitarRepositoryImpl(databaseCluster database.DatabaseCluster, redisCluster *redis.ClusterClient) repository.GuitarRepository {
	return &guitarRepositoryImpl{
		databaseCluster: databaseCluster,
		cache: cache.New(&cache.Options{
			Redis:      redisCluster,
			LocalCache: cache.NewTinyLFU(1000, time.Minute),
		}),
	}
}

type guitarRepositoryImpl struct {
	databaseCluster database.DatabaseCluster
	cache           *cache.Cache
}

func (g guitarRepositoryImpl) FindAll(ctx context.Context) ([]entity.Guitar, error) {
	guitars := make([]entity.Guitar, 0)

	rows, err := g.databaseCluster.Slave().QueryContext(ctx, "SELECT id, name, description, sku, price, image, type, strings, quantity_available, created_at, modified_at, deleted_at FROM guitar")
	if err != nil {
		return guitars, fmt.Errorf("error selecting guitars: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var guitar entity.Guitar
		if err = rows.Scan(&guitar.Id, &guitar.Name, &guitar.Description, &guitar.Sku, &guitar.Price, &guitar.Image, &guitar.Type, &guitar.Strings, &guitar.QuantityAvailable, &guitar.CreatedAt, &guitar.ModifiedAt, &guitar.DeletedAt); err != nil {
			return guitars, fmt.Errorf("error scanning row: %w", err)
		}
		guitars = append(guitars, guitar)
	}

	return guitars, nil
}

func (g guitarRepositoryImpl) FindById(ctx context.Context, id uuid.UUID) (guitar entity.Guitar, err error) {
	if err = g.cache.Get(ctx, id.String(), &guitar); err == nil {
		return guitar, nil
	}

	rows, err := g.databaseCluster.Slave().QueryContext(ctx, "SELECT id, name, description, sku, price, image, type, strings, quantity_available, created_at, modified_at, deleted_at FROM guitar WHERE id = $1", id)
	if err != nil {
		return guitar, fmt.Errorf("error selecting guitar: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&guitar.Id, &guitar.Name, &guitar.Description, &guitar.Sku, &guitar.Price, &guitar.Image, &guitar.Type, &guitar.Strings, &guitar.QuantityAvailable, &guitar.CreatedAt, &guitar.ModifiedAt, &guitar.DeletedAt); err != nil {
			return guitar, fmt.Errorf("error scanning row: %w", err)
		}
	}

	_ = g.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   id.String(),
		Value: guitar,
		TTL:   time.Hour,
	})

	return guitar, nil
}
