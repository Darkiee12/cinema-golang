package sdkredis

import (
	"cinema/component/appctx"
	"context"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"time"
)

type redisCache struct {
	store  *cache.Cache
	client *redis.Client
}

func NewRedisCache(appContext appctx.AppContext) *redisCache {
	rdClient := appContext.GetRedisClient()
	c := cache.New(&cache.Options{
		Redis:      rdClient,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	return &redisCache{store: c, client: rdClient}
}
func (rdc *redisCache) SetTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return rdc.store.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   ttl,
	})
}
func (rdc *redisCache) Set(ctx context.Context, key string, value interface{}) error {
	return rdc.store.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
	})

}
func (rdc *redisCache) Get(ctx context.Context, key string, value interface{}) error {
	return rdc.store.Get(ctx, key, &value)
}

func (rdc *redisCache) Delete(ctx context.Context, key string) error {

	return rdc.store.Delete(ctx, key)
}
func (rdc *redisCache) Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd {

	return rdc.client.Scan(ctx, cursor, match, count)
}
