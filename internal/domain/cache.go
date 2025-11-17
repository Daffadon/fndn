package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	cache_template "github.com/daffadon/fndn/internal/template/cache"
)

func InitInMemoryConfig(i infra.CommandRunner, path *string, inMemory *string) error {
	if path != nil {
		folderName := "/config/cache"
		var template, fileName string
		switch *inMemory {
		case "redis":
			fileName = folderName + "/redis.go"
			template = cache_template.RedisConfigTemplate
		case "valkey":
			fileName = folderName + "/valkey.go"
			template = cache_template.ValkeyConfigTemplate
		case "dragonfly":
			fileName = folderName + "/dragonfly.go"
			template = cache_template.DragonflyConfigTemplate
		case "redict":
			fileName = folderName + "/redict.go"
			template = cache_template.RedictConfigTemplate
		}
		if fileName != "" || template != "" {
			if err := pkg.GoFileGenerator(i, path, folderName, fileName, template); err != nil {
				log.Fatal(err)
				return err
			}
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitInMemoryConfigFile(i infra.CommandRunner, p *Project) error {
	if p.Path != nil {
		folderName := "/config/cache"
		var fileName, template string
		switch p.InMemory {
		case "valkey":
			fileName = folderName + "/valkey.acl"
			template = cache_template.ValkeyConfigFileTemplate
		}
		if fileName != "" || template != "" {
			if err := pkg.GenericFileGenerator(i, p.Path, folderName, fileName, template); err != nil {
				log.Fatal(err)
				return err
			}
		}
		return nil
	}
	return errors.New("path is nil")
}
