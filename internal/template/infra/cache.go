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
const ValkeyInfraTemplate string =`
package cache_infra

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/valkey-io/valkey-go"
	"github.com/valkey-io/valkey-go/valkeycompat"
)

type (
	ValkeyInfra interface {
		Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
		Get(ctx context.Context, key string) (string, error)
		Delete(ctx context.Context, key string) error
	}
	valkeInfra struct {
		valkeyClient valkeycompat.Cmdable
		logger       zerolog.Logger
	}
)

func NewValkeyCache(valkeyClient valkey.Client, logger zerolog.Logger) ValkeyInfra {
	compat := valkeycompat.NewAdapter(valkeyClient)
	if err := compat.Ping(context.Background()).Err(); err != nil {
		logger.Fatal().Err(err).Msg("failed to ping valkey client")
	}
	return &valkeInfra{
		valkeyClient: compat,
		logger:       logger,
	}
}

func (v *valkeInfra) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := v.valkeyClient.Set(ctx, key, value, expiration).Err()
	if err != nil {
		v.logger.Error().Err(err).Str("key", key).Msg("failed to set value in redis")
		return err
	}
	return nil
}

func (v *valkeInfra) Get(ctx context.Context, key string) (string, error) {
	val, err := v.valkeyClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			v.logger.Warn().Str("key", key).Msg("key not found in redis")
			return "", err
		}
		v.logger.Error().Err(err).Str("key", key).Msg("failed to get value from redis")
		return "", err
	}
	return val, nil
}

func (v *valkeInfra) Delete(ctx context.Context, key string) error {
	err := v.valkeyClient.Del(ctx, key).Err()
	if err != nil {
		v.logger.Error().Err(err).Str("key", key).Msg("failed to delete key from redis")
		return err
	}
	return nil
}
`
const DragonFlyInfraTemplate string =`
package cache_infra

import (
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type (
	DragonflyInfra interface {
		Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
		Get(ctx context.Context, key string) (string, error)
		Delete(ctx context.Context, key string) error
	}
	dragonflyInfra struct {
		dragonflyClient *redis.Client
		logger      zerolog.Logger
	}
)

func NewDragonflyCache(dragonflyClient *redis.Client, logger zerolog.Logger) DragonflyInfra {
	return &dragonflyInfra{
		dragonflyClient: dragonflyClient,
		logger:      logger,
	}
}

func (r *dragonflyInfra) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := r.dragonflyClient.Set(ctx, key, value, expiration).Err()
	if err != nil {
		r.logger.Error().Err(err).Str("key", key).Msg("failed to set value in dragonfly")
		return err
	}
	return nil
}

func (r *dragonflyInfra) Get(ctx context.Context, key string) (string, error) {
	val, err := r.dragonflyClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			r.logger.Warn().Str("key", key).Msg("key not found in dragonfly")
			return "", err
		}
		r.logger.Error().Err(err).Str("key", key).Msg("failed to get value from dragonfly")
		return "", err
	}
	return val, nil
}

func (r *dragonflyInfra) Delete(ctx context.Context, key string) error {
	err := r.dragonflyClient.Del(ctx, key).Err()
	if err != nil {
		r.logger.Error().Err(err).Str("key", key).Msg("failed to delete key from dragonfly")
		return err
	}
	return nil
}
`

const RedictInfraTemplate string =`
package cache_infra

import (
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type (
	RedictInfra interface {
		Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
		Get(ctx context.Context, key string) (string, error)
		Delete(ctx context.Context, key string) error
	}
	redictInfra struct {
		redictClient *redis.Client
		logger      zerolog.Logger
	}
)

func NewRedictCache(redictClient *redis.Client, logger zerolog.Logger) RedictInfra {
	return &redictInfra{
		redictClient: redictClient,
		logger:      logger,
	}
}

func (r *redictInfra) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := r.redictClient.Set(ctx, key, value, expiration).Err()
	if err != nil {
		r.logger.Error().Err(err).Str("key", key).Msg("failed to set value in redict")
		return err
	}
	return nil
}

func (r *redictInfra) Get(ctx context.Context, key string) (string, error) {
	val, err := r.redictClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			r.logger.Warn().Str("key", key).Msg("key not found in redict")
			return "", err
		}
		r.logger.Error().Err(err).Str("key", key).Msg("failed to get value from redict")
		return "", err
	}
	return val, nil
}

func (r *redictInfra) Delete(ctx context.Context, key string) error {
	err := r.redictClient.Del(ctx, key).Err()
	if err != nil {
		r.logger.Error().Err(err).Str("key", key).Msg("failed to delete key from redict")
		return err
	}
	return nil
}
`