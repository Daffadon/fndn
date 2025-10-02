package database_template

const MariaDBConfigTemplate string = `
package storage

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func NewSQLConn(logger zerolog.Logger)*sql.DB {
	host := viper.GetString("database.sql.host")
	user := viper.GetString("database.sql.user")
	password := viper.GetString("database.sql.password")
	port := viper.GetString("database.sql.port")
	dbname := viper.GetString("database.sql.name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		user, password, host, port, dbname)

	if dsn == "" {
		logger.Fatal().Msg("MariaDB configuration is not set")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to MariaDB")
	}

	if err := db.Ping(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to ping MariaDB")
	}
	return db
}
`

const DockerComposeMariaDBConfigTemplate string = `
# db
	{{.ProjectName}}_db:
		image: mariadb:12.0.2-ubi
		container_name: {{.ProjectName}}_db
		environment:
      MARIADB_ROOT_PASSWORD: ${DB_PASSWORD}
			MARIADB_USER: ${DB_USER}
			MARIADB_PASSWORD: ${DB_PASSWORD}
			MARIADB_DATABASE: ${DB_NAME}
		restart: unless-stopped
		ports:
			- "3306:3306"
		volumes:
			- {{.ProjectName}}_db_data:/var/lib/mysql
`
