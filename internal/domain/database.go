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
		fileName := folderName + "/db.go"
		var template string

		switch *db {
		case "postgresql":
			template = database_template.PostgresqlConfigTemplate
		case "mariadb":
			template = database_template.MariaDBConfigTemplate
		case "clickhouse":
			template = database_template.ClickHouseConfigTemplate
		case "mongodb":
			template = database_template.MongoDBConfigTemplate
		case "ferretdb":
			template = database_template.FerretDBConfigTemplate
		case "neo4j":
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

func GenerateSpecificDatabase(db string, infraRunner infra.CommandRunner, path string) error {
	// check folder config/router/ exist or not
	// check filename
	folderName := "/config/storage"
	fileName := folderName + "/db.go"

	// if exist, the file name add _framework_name
	exist := pkg.IsFileExists("." + fileName)
	if exist {
		fileName = folderName + "/db_" + db + ".go"
	}

	var template string
	switch db {
	case "postgresql":
		template = database_template.PostgresqlConfigTemplate
	case "mariadb":
		template = database_template.MariaDBConfigTemplate
	case "clickhouse":
		template = database_template.ClickHouseConfigTemplate
	case "mongodb":
		template = database_template.MongoDBConfigTemplate
	case "ferretdb":
		template = database_template.FerretDBConfigTemplate
	case "neo4j":
		template = database_template.Neo4jConfigTemplate
	}

	if err := pkg.GoFileGenerator(infraRunner, &path, folderName, fileName, template); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
