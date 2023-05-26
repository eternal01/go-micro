package cache

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/logx"
)

func GetCache(ctx context.Context, key string) any {
	c, err := collection.NewCache(0)
	if err != nil {
		logx.WithContext(ctx).Errorf("【CACHE-ERR】 %+v", err)
	}
	result, ok := c.Get(key)
	if !ok {
		logx.WithContext(ctx).Errorf("【CACHE-ERR】 %+v", err)
	}
	return result
}

func SetCache(ctx context.Context, key string, value any, expire time.Duration) {
	c, err := collection.NewCache(0)
	if err != nil {
		logx.WithContext(ctx).Errorf("【CACHE-ERR】 %+v", err)
	}
	c.SetWithExpire(key, value, expire)
}
