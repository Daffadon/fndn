package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	infra_template "github.com/daffadon/fndn/internal/template/infra"
)

func InitQuerierInfra(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := "/internal/infra/storage"
		fileName := folderName + "/querier.go"
		if err := pkg.GoFileGenerator(i, path, folderName, fileName, infra_template.QuerierInfraTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitRedisInfra(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := "/internal/infra/cache"
		fileName := folderName + "/redis.go"
		if err := pkg.GoFileGenerator(i, path, folderName, fileName, infra_template.RedisInfraTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitJetstreamInfra(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := "/internal/infra/mq"
		fileName := folderName + "/jetstream_infra.go"
		if err := pkg.GoFileGenerator(i, path, folderName, fileName, infra_template.JetstreamInfraTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitMinioInfra(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := "/internal/infra/storage"
		fileName := folderName + "/minio.go"
		if err := pkg.GoFileGenerator(i, path, folderName, fileName, infra_template.MinioInfraTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
