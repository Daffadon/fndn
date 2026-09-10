package main_template

const HTTPServerTemplate string = `
package server

import (
	{{.FrameworkImport}}
	{{if .HasDB}}{{.DBImport}}{{end}}
	{{if .HasMQ}}{{.MQImport}}{{end}}
	{{if .HasCache}}{{.CacheImport}}{{end}}
	"github.com/rs/zerolog"
	"go.uber.org/dig"
	"{{.ModuleName}}/internal/domain/handler"
)

type Server struct {
	Container   *dig.Container
	ServerReady chan bool
	Address     string
}

func (s *Server) Run(ctx context.Context) {
	err := s.Container.Invoke(
		func(
			logger zerolog.Logger,
			r {{.FrameworkRouter}},
			{{if .HasCache}}cache {{.CacheInstanceType}},{{end}}
			{{if .HasMQ}}{{.MQInstance}}{{end}}
			{{if .HasDB}}db {{.DBInstanceType}},{{end}}
			th handler.TodoHandler,
			// and many other returned type provided
			// in the container from /cmd/di/container.go
		) {
			{{if .HasCache}}{{.CacheCloseConn}}{{end}}
			{{if .HasMQ}}{{.MQCloseConn}}{{end}}
			{{if .HasDB}}defer {{.DBCloseConnection}}{{end}}
			
			// you can register your routes here
			// for the example and implementation, here is the example

			handler.RegisterTodoRoutes(r,th)
			
			srv := &http.Server{
				Addr:              s.Address,
				Handler:           {{.RouterHandler}},
				ReadHeaderTimeout: 5 * time.Second,
			}
			go func() {
				e := os.Getenv("ENV")
				switch e {
				case "production":
					if err := srv.ListenAndServeTLS("./config/cert/server.crt", "./config/cert/server.key"); err != nil && err != http.ErrServerClosed {
						logger.Fatal().Err(err).Msg("Failed to listen and serve http server")
					}
				default:
					if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
						logger.Fatal().Err(err).Msg("Failed to listen and serve http server")
					}
				}
			}()

			if s.ServerReady != nil {
				for range 50 {
					conn, err := net.DialTimeout("tcp", s.Address, 100*time.Millisecond)
					if err == nil {
						if err := conn.Close(); err != nil {
							logger.Fatal().Err(err).Msg("establish check connection failed to close")
						}
						s.ServerReady <- true
						break
					}
					time.Sleep(100 * time.Millisecond)
				}
			}

			logger.Info().Msgf("HTTP Server Starting in port %s", s.Address)
			<-ctx.Done()

			logger.Info().Msg("Shutting down server...")

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				logger.Fatal().Err(err).Msg("HTTP Server forced to shutdown")
			}

			logger.Info().Msg("Server exiting...")
		})
	if err !=nil{
		log.Fatalf("failed to initialize application: %v", err)
	}
}
`
