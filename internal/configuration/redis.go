package configuration

import (
	"github.com/redis/go-redis/v9"
	"strings"
)

func NewRedisCluster(redisAddrs string) *redis.ClusterClient {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: strings.Split(redisAddrs, ","),
	})

	return client
}
