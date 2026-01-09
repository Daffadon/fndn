package database_template

const FerretDBConfigTemplate string = `
package storage

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewFerretDBConn(logger zerolog.Logger) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	protocol := viper.GetString("database.nosql.protocol")
	host := viper.GetString("database.nosql.host")
	port := viper.GetString("database.nosql.port")

	uri := fmt.Sprintf("%s://%s:%s/", protocol, host, port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logger.Fatal().Err(err).Msg("FerretDB configuration is not set")
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Fatal().Err(err).Msg("Failed to ping FerretDB")
	}
	return client
}
`

const DockerComposeFerretDBConfigTemplate string = `
# db
	{{.ProjectName}}_db:
			image: ghcr.io/ferretdb/ferretdb:2.5.0
			container_name: {{.ProjectName}}_db
			restart: unless-stopped
			environment:
				- FERRETDB_POSTGRESQL_URL=postgres://${DB_USER}:${DB_PASSWORD}@postgres:5432/postgres
			ports:
				- "27017:27017"
			depends_on:
				- postgres 

	postgres:
    image: ghcr.io/ferretdb/postgres-documentdb:17-0.106.0-ferretdb-2.5.0
		restart: unless-stopped
    environment:
      - POSTGRES_USER=${DB_USER}
      - POSTGRES_PASSWORD=${DB_PASSWORD}
      - POSTGRES_DB=postgres
    volumes:
			- {{.ProjectName}}_db_data:/data/db
`
