package main_template

const HTTPServerTemplate string = `
package server

import "github.com/redis/go-redis/v9"

type Server struct {
	Container   *dig.Container
	ServerReady chan bool
	Address     string
}

func (s *Server) Run(ctx context.Context) {
	err := s.Container.Invoke(
		func(
			logger zerolog.Logger,
			router *gin.Engine,
			redis *redis.Client,
			nc *nats.Conn,
			pgx *pgxpool.Pool,
			// and many other returned type provided
			// in the container from /cmd/di/container.go
		) {
			defer func() {
				if err := redis.Close(); err != nil {
					logger.Error().Err(err).Msg("Failed to close Redis client")
				}
			}()
			defer func() {
				if err := nc.Drain(); err != nil {
					logger.Error().Err(err).Msg("Failed to drain nats client")
				}
			}()
			defer pgx.Close()

			router.Use(gin.Recovery())
			srv := &http.Server{
				Addr:              s.Address,
				Handler:           router,
				ReadHeaderTimeout: 5 * time.Second,
			}
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Fatal().Err(err).Msg("Failed to listen and server http server")
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
