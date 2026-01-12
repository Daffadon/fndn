package framework_template

const GorillaMuxConfigTemplate string = `
package router

import (

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func NewHTTPMux() *mux.Router {
	r := mux.NewRouter()
	return r
}

func WarpWithCorsAndLogger(r *mux.Router) http.Handler {
	allowOrigins := viper.GetString("server.cors.allow_origins")
	allowMethods := viper.GetString("server.cors.allow_methods")
	allowHeaders := viper.GetString("server.cors.allow_headers")
	exposeHeaders := viper.GetString("server.cors.expose_headers")
	allowCredentials := viper.GetBool("server.cors.allow_credential")
	maxAge := viper.GetInt("server.cors.max_age")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{allowOrigins},
		AllowedMethods:   []string{allowMethods},
		AllowedHeaders:   []string{allowHeaders},
		ExposedHeaders:   []string{exposeHeaders},
		AllowCredentials: allowCredentials,
		MaxAge:           maxAge,
	})

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	return c.Handler(loggedRouter)
}
`
