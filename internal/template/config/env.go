package config_template

const ENVConfigTemplate string = `
package env

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
