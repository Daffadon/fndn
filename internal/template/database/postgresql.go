package database_template

const PostgresqlConfigTemplate string = `
package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func NewSQLConn(logger zerolog.Logger) *pgxpool.Pool {
	protocol := viper.GetString("database.sql.protocol")
	host := viper.GetString("database.sql.host")
	user := viper.GetString("database.sql.user")
	password := viper.GetString("database.sql.password")
	port := viper.GetString("database.sql.port")
	dbname := viper.GetString("database.sql.name")
	sslmode := viper.GetString("database.sql.sslmode")
	
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s", protocol, user, password, host, port, dbname, sslmode)
	if dsn == "" {
		logger.Fatal().Msg("Database configuration is not set")
	}
	
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		logger.Fatal().Err(err).Msg("Database configuration is not set")
	}
	
	if err := pool.Ping(context.Background()); err != nil {
			logger.Fatal().Err(err).Msg("Failed to ping PostgreSQL")
	}
	return pool
}
`
const DockerComposePostgresqlConfigTemplate string = `
# db
  {{.ProjectName}}_db:
    image: postgres:17.5-alpine3.22
    container_name: {{.ProjectName}}_db
    environment:
      POSTGRES_USER: ${DB_USER}
      POSTGRES_PASSWORD: ${DB_PASSWORD}
      POSTGRES_DB: ${DB_NAME}
    restart: unless-stopped
    ports:
      - "5432:5432"
    volumes:
      - {{.ProjectName}}_db_data:/var/lib/postgresql/data
`
