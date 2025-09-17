package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	database_template "github.com/daffadon/fndn/internal/template/database"
)

func InitPostgresqlConfig(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := "/config/storage"
		fileName := folderName + "/postgresql.go"
		if err := pkg.FileGenerator(i, path, folderName, fileName, database_template.PostgresqlConfigTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
