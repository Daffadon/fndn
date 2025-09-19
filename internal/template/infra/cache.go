package infra_template

const RedisInfraTemplate string =`
package cache_infra

import 	(
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type (
	RedisInfra interface {
		Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
		Get(ctx context.Context, key string) (string, error)
		Delete(ctx context.Context, key string) error
	}
	redisInfra struct {
		redisClient *redis.Client
		logger      zerolog.Logger
	}
)

func NewRedisCache(redisClient *redis.Client, logger zerolog.Logger) RedisInfra {
	return &redisInfra{
		redisClient: redisClient,
		logger:      logger,
	}
}

func (r *redisInfra) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := r.redisClient.Set(ctx, key, value, expiration).Err()
	if err != nil {
		r.logger.Error().Err(err).Str("key", key).Msg("failed to set value in redis")
		return err
	}
	return nil
}

func (r *redisInfra) Get(ctx context.Context, key string) (string, error) {
	val, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			r.logger.Warn().Str("key", key).Msg("key not found in redis")
			return "", err
		}
		r.logger.Error().Err(err).Str("key", key).Msg("failed to get value from redis")
		return "", err
	}
	return val, nil
}

func (r *redisInfra) Delete(ctx context.Context, key string) error {
	err := r.redisClient.Del(ctx, key).Err()
	if err != nil {
		r.logger.Error().Err(err).Str("key", key).Msg("failed to delete key from redis")
		return err
	}
	return nil
}
`