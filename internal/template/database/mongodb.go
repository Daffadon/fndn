package database_template

const MongoDBConfigTemplate string = `
package storage

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)
	
func NewMongoDBConn(logger zerolog.Logger) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	protocol := viper.GetString("database.nosql.protocol")
	host := viper.GetString("database.nosql.host")
	user := viper.GetString("database.nosql.user")
	password := viper.GetString("database.nosql.password")
	port := viper.GetString("database.nosql.port")
	dbname := viper.GetString("database.nosql.name")

	uri := fmt.Sprintf("%s://%s:%s@%s:%s/%s?authSource=admin", protocol, user, password, host, port, dbname)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logger.Fatal().Err(err).Msg("Database configuration is not set")
	}
	// Ping the database to verify connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Fatal().Err(err).Msg("Failed to ping MongoDB")
	}
	return client
}
`

const DockerComposeMongoDBConfigTemplate string = `
# db
	{{.ProjectName}}_db:
			image: mongo:8.0-noble
			container_name: {{.ProjectName}}_db
			restart: unless-stopped
			environment:
				MONGO_INITDB_ROOT_USERNAME: ${DB_USER}
				MONGO_INITDB_ROOT_PASSWORD: ${DB_PASSWORD}
				MONGO_INITDB_DATABASE: ${DB_NAME}
			ports:
				- "27017:27017"
			volumes:
				- {{.ProjectName}}_db_data:/data/db
`
