package domain

import (
	"errors"
	"log"
	"sync"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	cache_template "github.com/daffadon/fndn/internal/template/cache"
	config_template "github.com/daffadon/fndn/internal/template/config"
	database_template "github.com/daffadon/fndn/internal/template/database"
	mq_template "github.com/daffadon/fndn/internal/template/mq"
	objectstorage_template "github.com/daffadon/fndn/internal/template/object_storage"
)

func InitDockerFileConfig(i infra.CommandRunner, path *string, projectName string) error {
	if path != nil {
		folderName := ""
		fileName := folderName + "/Dockerfile"
		st := struct {
			ProjectName string
		}{
			ProjectName: projectName,
		}
		c, err := pkg.ParseTemplate(config_template.DockerfileConfigTemplate, st)
		if err != nil {
			log.Fatal(err)
			return err
		}
		if err := pkg.GenericFileGenerator(i, path, folderName, fileName, c); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitDockerComposeConfig(i infra.CommandRunner, path *string, projectName string) error {
	if path != nil {
		folderName := ""
		fileName := folderName + "/docker-compose.yml"
		st := struct {
			ProjectName string
		}{
			ProjectName: projectName,
		}
		var results []string
		templates := []string{
			config_template.DockerComposeAppConfigTemplate,
			database_template.DockerComposePostgresqlConfigTemplate,
			mq_template.DockerComposeNatsConfigTemplate,
			cache_template.DockerComposeRedisConfigTemplate,
			objectstorage_template.DockerComposeMinioConfigTemplate,
			database_template.DockerComposePostgresqlVolumeTemplate,
			mq_template.DockerComposeNatsVolumeTemplate,
			cache_template.DockerComposeRedisVolumeTemplate,
			objectstorage_template.DockerComposeMinioVolumeTemplate,
		}
		for _, tpl := range templates {
			if err := parserHelper(&results, tpl, st); err != nil {
				log.Fatal(err)
				return err
			}
		}

		s := config_template.DockerComposeDefaultConfigTemplate
		for i := range results {
			if i == 5 {
				s += config_template.DockerComposeVolumeConfigTemplate
			}
			s += results[i]
		}
		if err := pkg.GenericFileGenerator(i, path, folderName, fileName, s); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func parserHelper(results *[]string, template string, st interface{}) error {
	var mu sync.Mutex
	c, err := pkg.ParseTemplate(template, st)
	if err != nil {
		return err
	}
	mu.Lock()
	*results = append(*results, c)
	mu.Unlock()
	return nil
}
