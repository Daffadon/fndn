package cache_template

const RedisConfigTemplate string = `
package cache

import (
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
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

const DockerComposeRedisConfigTemplate string = `
# redis
  {{.ProjectName}}_cache:
    image: redis:7.2
    container_name: {{.ProjectName}}_cache
    volumes:
      - {{.ProjectName}}_redis_data:/data
    restart: unless-stopped
		ports:
      - "6379:6379"
    environment:
      - REDIS_PASSWORD=${REDIS_PASSWORD}
    command: ["redis-server", "--requirepass", "$REDIS_PASSWORD"]
`

const DockerComposeRedisVolumeTemplate string = `
  {{.ProjectName}}_redis_data: {}`
