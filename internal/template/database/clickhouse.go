package database_template

const ClickHouseConfigTemplate string = `
package storage

import (
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func NewClickhouseConn(logger zerolog.Logger) clickhouse.Conn {
	addr := fmt.Sprintf("%s:%s", viper.GetString("database.sql.host"), viper.GetString("database.sql.port"))
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{addr},
		Auth: clickhouse.Auth{
			Database: viper.GetString("database.sql.name"),
			Username: viper.GetString("database.sql.user"),
			Password: viper.GetString("database.sql.password"),
		},
		TLS:         nil,
		DialTimeout: 5_000_000_000,
		Debug:       viper.GetBool("database.sql.debug"),
	})
	if err != nil {
		logger.Fatal().Msgf("failed to open ClickHouse connection: %v", err)
	}
	if err := conn.Ping(context.Background()); err != nil {
		log.Fatalf("failed to ping ClickHouse: %v", err)
	}
	return conn
}
`

const DockerComposeClickHouseConfigTemplate string = `
# db
	{{.ProjectName}}_db:
			image: dhi.io/clickhouse-server:25.12
			container_name: {{.ProjectName}}_db
			restart: unless-stopped
			environment:
				CLICKHOUSE_DB: ${DB_NAME}
				CLICKHOUSE_USER: ${DB_USER}
				CLICKHOUSE_PASSWORD: ${DB_PASSWORD}
			ports:
				- "9002:9000"
			volumes:
			- {{.ProjectName}}_db_data:/var/lib/clickhouse
`
