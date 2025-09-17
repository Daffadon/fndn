package cache_template

const RedisConfigTemplate string = `
package cache

import (
	"github.com/redis/go-redis/v9"
)

func NewRedisConnection(zerolog zerolog.Logger) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", viper.GetString("redis.address"), viper.GetString("redis.port"))
	client := redis.NewClient(&redis.Options{
		Addr:       addr,
		ClientName: viper.GetString("redis.client_name"),
		Protocol:   2,
		Password:   viper.GetString("redis.password"),
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		zerolog.Fatal().Err(err).Msg("failed to connect to redis")
	}
	return client, nil
}
`
