package database_template

const PostgresqlConfigTemplate string = `
package storage

import "github.com/jackc/pgx/v5/pgxpool"

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
	return pool
}
`
