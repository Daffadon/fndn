package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	objectstorage_template "github.com/daffadon/fndn/internal/template/object_storage"
)

func InitMinioConfig(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := "/config/storage"
		fileName := folderName + "/minio.go"
		if err := pkg.GoFileGenerator(i, path, folderName, fileName, objectstorage_template.MinioConfigTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
