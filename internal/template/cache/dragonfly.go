package cache_template

const DragonflyConfigTemplate string = `
package cache
import (
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func NewDragonflyConnection(zerolog zerolog.Logger) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", viper.GetString("dragonfly.address"), viper.GetString("dragonfly.port"))
	client := redis.NewClient(&redis.Options{
		Addr:       addr,
		ClientName: viper.GetString("dragonfly.client_name"),
		Protocol:   2,
		Password:   viper.GetString("dragonfly.password"),
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		zerolog.Fatal().Err(err).Msg("failed to connect to dragonfly")
	}
	return client, nil
}
`

const DockerComposeDragonflyConfigTemplate string = `
# dragonfly
  {{.ProjectName}}_cache:
    image: ghcr.io/dragonflydb/dragonfly:v1.34.1
    container_name: {{.ProjectName}}_cache
    volumes:
      - {{.ProjectName}}_dragonfly_data:/data
    restart: unless-stopped
		ports:
      - "6379:6379"
    environment:
      - CACHE_PASSWORD=${CACHE_PASSWORD}
    command: ["dragonfly", "--requirepass", "$CACHE_PASSWORD", "--proactor_threads", "2"]
`

const DockerComposeDragonflyVolumeTemplate string = `
  {{.ProjectName}}_dragonfly_data: {}`
