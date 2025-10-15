package config_template

const ENVConfigTemplate string = `
package env

	import 	"github.com/spf13/viper"

	// config.yaml that loaded by this code in production environment (container)
	// can be achieved by mount the config file to the container at / (default) or read the config file
	// from remote source, which can you do based on docs (https://github.com/spf13/viper)

	// or (which i don't suggest but working) inject the config.yaml in the Dockerfile.

func Load() {
	env := os.Getenv("ENV")
	configName := "config.local"
	configpath := "."
	switch env {
	case "production":
		configName = "config"
	case "test":
		configName = "config.test"
	}

	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configpath)
	if err := viper.ReadInConfig(); err != nil {
		panic("failed to read config")
	}
}
`

const DotENVExampleTemplate string = `
# copy this file and change name to .env and fill the variable
# to be used by docker compose

# app env
ENV=

# db env
DB_USER=
DB_PASSWORD=
DB_NAME=

# nats env
MQ_USER=
MQ_PASSWORD=

# redis env
REDIS_PASSWORD=

# minio env
MINIO_ROOT_USER=
MINIO_ROOT_PASSWORD=
`
