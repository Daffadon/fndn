package framework_template

const ChiConfigTemplate string = `
package router 

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/spf13/viper"
)

func NewHTTP() *chi.Mux{
	r := chi.NewRouter()
	env := os.Getenv("ENV")
	if env != "production" {
		r.Use(middleware.Logger)
	}

	allowOrigins := viper.GetString("server.cors.allow_origins")
	allowMethods := viper.GetString("server.cors.allow_methods")
	allowHeaders := viper.GetString("server.cors.allow_headers")
	exposeHeaders := viper.GetString("server.cors.expose_headers")
	allowCredentials := viper.GetBool("server.cors.allow_credential")
	maxAge := viper.GetInt("server.cors.max_age")

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   strings.Split(allowOrigins, ","),
		AllowedMethods:   strings.Split(allowMethods, ","),
		AllowedHeaders:   strings.Split(allowHeaders, ","),
		ExposedHeaders:   strings.Split(exposeHeaders, ","),
		AllowCredentials: allowCredentials,
		MaxAge:           maxAge,
	}).Handler)

	return r
}
`
