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

func InitDockerComposeConfig(i infra.CommandRunner, p *Project) error {
	if p.Path != nil {
		folderName := ""
		fileName := folderName + "/docker-compose.yml"
		st := struct {
			ProjectName string
		}{
			ProjectName: p.Name,
		}
		var results []string
		var dbDockerTemplate string

		switch p.Database {
		case "postgresql":
			dbDockerTemplate = database_template.DockerComposePostgresqlConfigTemplate
		case "mariadb":
			dbDockerTemplate = database_template.DockerComposeMariaDBConfigTemplate
		case "mongodb":
			dbDockerTemplate = database_template.DockerComposeMongoDBConfigTemplate
		case "ferretdb":
			dbDockerTemplate = database_template.DockerComposeFerretDBConfigTemplate
		case "neo4j":
			dbDockerTemplate = database_template.DockerComposeNeo4JConfigTemplate
		}
		templates := []string{
			config_template.DockerComposeAppConfigTemplate,
			dbDockerTemplate,
			mq_template.DockerComposeNatsConfigTemplate,
			cache_template.DockerComposeRedisConfigTemplate,
			objectstorage_template.DockerComposeMinioConfigTemplate,

			// volume
			database_template.DockerComposeDBVolumeTemplate,
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
		if err := pkg.GenericFileGenerator(i, p.Path, folderName, fileName, s); err != nil {
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
