package framework_template

const EchoConfigTemplate string = `
package router 

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func NewHTTPEcho() *echo.Echo {
	env := os.Getenv("ENV")
	e := echo.New()
	e.Use(middleware.RequestLogger())

	if env == "production" {
		e.Debug = false
		e.HideBanner = true
		e.HidePort = true
	}

	allowOrigins := viper.GetString("server.cors.allow_origins")
	allowMethods := viper.GetString("server.cors.allow_methods")
	allowHeaders := viper.GetString("server.cors.allow_headers")
	exposeHeaders := viper.GetString("server.cors.expose_headers")
	allowCredentials := viper.GetBool("server.cors.allow_credential")
	maxAge := viper.GetInt("server.cors.max_age")

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     strings.Split(allowOrigins, ","),
		AllowMethods:     strings.Split(allowMethods, ","),
		AllowHeaders:     strings.Split(allowHeaders, ","),
		ExposeHeaders:    strings.Split(exposeHeaders, ","),
		AllowCredentials: allowCredentials,
		MaxAge:           int(time.Duration(maxAge).Seconds()),
	}))

	return e
}
`
