package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	cache_template "github.com/daffadon/fndn/internal/template/cache"
)

func InitRedisConfig(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := "/config/cache"
		fileName := folderName + "/redis.go"
		if err := pkg.FileGenerator(i, path, folderName, fileName, cache_template.RedisConfigTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
