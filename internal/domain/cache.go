package domain

import (
	"errors"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	cache_template "github.com/daffadon/fndn/internal/template/cache"
)

var cacheTemplates = map[string]string{
	"redis":     cache_template.RedisConfigTemplate,
	"valkey":    cache_template.ValkeyConfigTemplate,
	"dragonfly": cache_template.DragonflyConfigTemplate,
	"redict":    cache_template.RedictConfigTemplate,
}

func InitInMemoryConfig(path *string, inMemory *string) error {
	if path != nil {
		folderName := "/config/cache"
		fileName := folderName + "/cache.go"

		template, err := lookup(cacheTemplates, "cache", *inMemory)
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

func InitInMemoryConfigFile(p *Project) error {
	if p.Path != nil {
		folderName := "/config/cache"
		var fileName, template string
		switch p.InMemory {
		case "valkey":
			fileName = folderName + "/valkey.acl"
			template = cache_template.ValkeyConfigFileTemplate
		default:
			return nil
		}
		if err := pkg.GenericFileGenerator(p.Path, folderName, fileName, template); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func GenerateSpecificCache(cache string, path string) error {
	folderName := "/config/cache"
	fileName := resolveTarget(folderName+"/cache.go", folderName+"/cache", cache)

	t, err := lookup(cacheTemplates, "cache", cache)
	if err != nil {
		return err
	}
	if err := pkg.GoFileGenerator(&path, folderName, fileName, t); err != nil {
		return err
	}
	return nil
}

// GenerateSpecificCachce keeps old typo'd name working.
func GenerateSpecificCachce(cache string, _ infra.CommandRunner, path string) error {
	return GenerateSpecificCache(cache, path)
}
