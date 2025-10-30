package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	database_template "github.com/daffadon/fndn/internal/template/database"
)

func InitDBConfig(i infra.CommandRunner, path *string, db *string) error {
	if path != nil {
		folderName := "/config/storage"
		var fileName, template string
		switch *db {
		case "postgresql":
			fileName = folderName + "/postgresql.go"
			template = database_template.PostgresqlConfigTemplate
		case "mariadb":
			fileName = folderName + "/mariadb.go"
			template = database_template.MariaDBConfigTemplate
		case "clickhouse":
			fileName = folderName + "/clickhouse.go"
			template = database_template.ClickHouseConfigTemplate
		case "mongodb":
			fileName = folderName + "/mongodb.go"
			template = database_template.MongoDBConfigTemplate
		case "ferretdb":
			fileName = folderName + "/ferretdb.go"
			template = database_template.FerretDBConfigTemplate
		case "neo4j":
			fileName = folderName + "/neo4j.go"
			template = database_template.Neo4jConfigTemplate
		}
		if err := pkg.GoFileGenerator(i, path, folderName, fileName, template); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
