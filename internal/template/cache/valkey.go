package cache_template

const ValkeyConfigTemplate string = `
package cache

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"github.com/valkey-io/valkey-go"
)

func NewValkeyConnection(logger zerolog.Logger) (valkey.Client, error) {
	addr := fmt.Sprintf("%s:%s", viper.GetString("valkey.address"), viper.GetString("valkey.port"))
	client, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{addr},
		Username:    viper.GetString("valkey.username"),
		Password:    viper.GetString("valkey.password"),
	})
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect to valkey")
	}
	return client, nil
}
`

const DockerComposeValkeyConfigTemplate string = `
# valkey
  {{.ProjectName}}_cache:
    image: valkey/valkey:9-alpine3.22
    container_name: {{.ProjectName}}_cache
    volumes:
      - {{.ProjectName}}_valkey_data:/data
			- ./config/cache/valkey.acl:/etc/valkey/valkey.acl
    restart: unless-stopped
		ports:
      - "6379:6379"
		command: ["valkey-server", "--aclfile", "/etc/valkey/valkey.acl"]
`

const DockerComposeValkeyVolumeTemplate string = `
  {{.ProjectName}}_valkey_data: {}`

const ValkeyConfigFileTemplate string = `user default off
user username on >password allcommands allkeys
`
