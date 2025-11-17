package cache_template

const RedictConfigTemplate = `
package cache

import (
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func NewRedictConnection(zerolog zerolog.Logger) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", viper.GetString("redict.address"), viper.GetString("redict.port"))
	client := redis.NewClient(&redis.Options{
		Addr:       addr,
		ClientName: viper.GetString("redict.client_name"),
		Protocol:   2,
		Password:   viper.GetString("redict.password"),
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		zerolog.Fatal().Err(err).Msg("failed to connect to redis")
	}
	return client, nil
}
`
const DockerComposeRedictConfigTemplate = `
# redict
	{{.ProjectName}}_cache:
		image: registry.redict.io/redict:7.3.6-scratch
		container_name: {{.ProjectName}}_cache
		volumes:
			- {{.ProjectName}}_redict_data:/data
		restart: unless-stopped
		ports:
			- "6379:6379"
		environment:
			- CACHE_PASSWORD=${CACHE_PASSWORD}
		command:
			[
				"redict-server",
				"--requirepass",
				"$CACHE_PASSWORD",
				"--maxmemory",
				"512mb",
				"--maxmemory-policy",
				"allkeys-lru",
			]
`
const DockerComposeRedictVolumeTemplate string = `
  {{.ProjectName}}_redict_data: {}`
