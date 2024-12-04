package impl

import (
	"context"
	"fmt"
	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
	"guitarStore/internal/configuration/database"
	"guitarStore/internal/entity"
	"guitarStore/internal/repository"
	"time"
)

func NewPromocodeRepositoryImpl(databaseCluster database.DatabaseCluster, redisCluster *redis.ClusterClient) repository.PromocodeRepository {
	return &promocodeRepositoryImpl{
		databaseCluster: databaseCluster,
		cache: cache.New(&cache.Options{
			Redis:      redisCluster,
			LocalCache: cache.NewTinyLFU(1000, time.Minute),
		}),
	}
}

type promocodeRepositoryImpl struct {
	databaseCluster database.DatabaseCluster
	cache           *cache.Cache
}

func (p promocodeRepositoryImpl) FindByCode(ctx context.Context, code string) (promocode entity.Promocode, err error) {
	if err = p.cache.Get(ctx, code, &promocode); err == nil {
		return promocode, nil
	}

	rows, err := p.databaseCluster.Slave().QueryContext(ctx, "SELECT id, name, description, code, max_usage, discount_amount, expired_at, created_at, modified_at, deleted_at FROM promocode WHERE code = $1", code)
	if err != nil {
		return promocode, fmt.Errorf("error selecting promocode: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&promocode.Id, &promocode.Name, &promocode.Description, &promocode.Code, &promocode.MaxUsage, &promocode.DiscountAmount, &promocode.ExpiredAt, &promocode.CreatedAt, &promocode.ModifiedAt, &promocode.DeletedAt); err != nil {
			return promocode, fmt.Errorf("error scanning row: %w", err)
		}
	}

	_ = p.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   code,
		Value: promocode,
		TTL:   time.Hour,
	})

	return promocode, nil
}
