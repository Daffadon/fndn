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
		fileName := folderName + "/cache.go"

		var template string
		switch *inMemory {
		case "redis":
			template = cache_template.RedisConfigTemplate
		case "valkey":
			template = cache_template.ValkeyConfigTemplate
		case "dragonfly":
			template = cache_template.DragonflyConfigTemplate
		case "redict":
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

func GenerateSpecificCachce(cache string, infraRunner infra.CommandRunner, path string) error {
	// check folder config/router/ exist or not
	// check filename
	folderName := "/config/cache"
	fileName := folderName + "/cache.go"

	// if exist, the file name add _framework_name
	exist := pkg.IsFileExists("." + fileName)
	if exist {
		fileName = folderName + "/cache_" + cache + ".go"
	}

	var t string
	switch cache {
	case "redis":
		t = cache_template.RedisConfigTemplate
	case "valkey":
		t = cache_template.ValkeyConfigTemplate
	case "dragonfly":
		t = cache_template.DragonflyConfigTemplate
	case "redict":
		t = cache_template.RedictConfigTemplate
	}

	if err := pkg.GoFileGenerator(infraRunner, &path, folderName, fileName, t); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
