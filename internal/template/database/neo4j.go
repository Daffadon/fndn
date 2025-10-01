package database_template

const Neo4jConfigTemplate string = `
package storage

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func NewGraphDBConn(logger zerolog.Logger) neo4j.DriverWithContext {
	protocol := viper.GetString("database.graph.protocol")
	host := viper.GetString("database.graph.host")
	port := viper.GetString("database.graph.port")
	username := viper.GetString("database.graph.user")
	password := viper.GetString("database.graph.password")

	if  username == "" || password == "" {
		logger.Fatal().Msg("Neo4j configuration is not set")
	}

	uri := fmt.Sprintf("%s://%s:%s",protocol,host,port)

	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to create Neo4j driver")
	}

	ctx := context.Background()
	if err := driver.VerifyConnectivity(ctx); err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to Neo4j")
	}

	return driver
}
`

const DockerComposeNeo4JConfigTemplate string = `
# db
  {{.ProjectName}}_db:
    image: neo4j:5.26.12-community-bullseye
    container_name: {{.ProjectName}}_db
    environment:
      NEO4J_AUTH: ${DB_USER}/${DB_PASSWORD}
      NEO4J_dbms_default__database: ${DB_NAME}
    restart: unless-stopped
    ports:
      - "7474:7474"
      - "7687:7687"
    volumes:
      - {{.ProjectName}}_db_data:/data
`
