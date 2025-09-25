package framework_template

const FiberConfigTemplate string = `
package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func NewHTTP() *fiber.App {
	env := os.Getenv("ENV")
	var config fiber.Config
	if env == "production" {
		config.Prefork = true
	}
	r := fiber.New(config)

	allowOrigins := viper.GetString("server.cors.allow_origins")
	allowMethods := viper.GetString("server.cors.allow_methods")
	allowHeaders := viper.GetString("server.cors.allow_headers")
	exposeHeaders := viper.GetString("server.cors.expose_headers")
	allowCredentials := viper.GetBool("server.cors.allow_credential")
	maxAge := viper.GetInt("server.cors.max_age")

	r.Use(cors.New(cors.Config{
			AllowOrigins:     allowOrigins,
			AllowMethods:     allowMethods,
			AllowHeaders:     allowHeaders,
			ExposeHeaders:    exposeHeaders,
			AllowCredentials: allowCredentials,
			MaxAge:           maxAge,
	}))

	return r
}
`
