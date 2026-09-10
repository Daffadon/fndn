package domain

import (
	"errors"

	"github.com/daffadon/fndn/internal/pkg"
	database_template "github.com/daffadon/fndn/internal/template/database"
)

var databaseTemplates = map[string]string{
	"postgresql": database_template.PostgresqlConfigTemplate,
	"mariadb":    database_template.MariaDBConfigTemplate,
	"clickhouse": database_template.ClickHouseConfigTemplate,
	"mongodb":    database_template.MongoDBConfigTemplate,
	"ferretdb":   database_template.FerretDBConfigTemplate,
	"neo4j":      database_template.Neo4jConfigTemplate,
}

func InitDBConfig(path *string, db *string) error {
	if *db == None {
		return nil
	}
	if path != nil {
		folderName := "/config/storage"
		fileName := folderName + "/db.go"

		template, err := lookup(databaseTemplates, "database", *db)
		if err != nil {
			return err
		}
		if err := pkg.GoFileGenerator(path, folderName, fileName, template); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func GenerateSpecificDatabase(db string, path string) error {
	folderName := "/config/storage"
	fileName := resolveTarget(folderName+"/db.go", folderName+"/db", db)

	template, err := lookup(databaseTemplates, "database", db)
	if err != nil {
		return err
	}
	if err := pkg.GoFileGenerator(&path, folderName, fileName, template); err != nil {
		return err
	}
	return nil
}
