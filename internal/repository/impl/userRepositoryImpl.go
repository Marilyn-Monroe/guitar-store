package impl

import (
	"github.com/redis/go-redis/v9"
	"guitarStore/internal/configuration/database"
	"guitarStore/internal/repository"
)

func NewUserRepositoryImpl(databaseCluster database.DatabaseCluster, redisCluster *redis.ClusterClient) repository.UserRepository {
	return &userRepositoryImpl{
		databaseCluster: databaseCluster,
		redisCluster:    redisCluster,
	}
}

type userRepositoryImpl struct {
	databaseCluster database.DatabaseCluster
	redisCluster    *redis.ClusterClient
}
