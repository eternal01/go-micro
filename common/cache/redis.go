package cache

import (
	"sync"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

var (
	Once           sync.Once
	RedisSingleton *Singleton
)

type Singleton struct {
	RedisCluster *redis.Redis
}

func GetSingleton(conf redis.RedisConf) *Singleton {
	Once.Do(func() {
		RedisSingleton = &Singleton{
			RedisCluster: redis.MustNewRedis(conf),
		}
	})
	return RedisSingleton
}

func (r *Singleton) GetRedisCache(key string) (string, error) {
	return r.RedisCluster.Get(key)
}

func (r *Singleton) SetRedisCache(key string, value string) error {
	return r.RedisCluster.Set(key, value)
}

func (r *Singleton) DelRedisCache(key string) (int, error) {
	return r.RedisCluster.Del(key)
}
