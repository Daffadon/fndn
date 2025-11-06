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
		case "clickhouse":
			dbDockerTemplate = database_template.DockerComposeClickHouseConfigTemplate
		case "mongodb":
			dbDockerTemplate = database_template.DockerComposeMongoDBConfigTemplate
		case "ferretdb":
			dbDockerTemplate = database_template.DockerComposeFerretDBConfigTemplate
		case "neo4j":
			dbDockerTemplate = database_template.DockerComposeNeo4JConfigTemplate
		}

		var mqDockerTemplate string
		var mqVolumetemplate string
		switch p.MQ {
		case "nats":
			mqDockerTemplate = mq_template.DockerComposeNatsConfigTemplate
			mqVolumetemplate = mq_template.DockerComposeNatsVolumeTemplate
		case "rabbitmq":
			mqDockerTemplate = mq_template.DockerComposeRabbitMQConfigTemplate
			mqVolumetemplate = mq_template.DockerComposeRabbitVolumeTemplate
		case "kafka":
			mqDockerTemplate = mq_template.DockerComposeKafkaConfigTemplate
			mqVolumetemplate = mq_template.DockerComposeKafkaVolumeTemplate
		}

		var cacheDockerTemplate string
		var cacheVolumeTemplate string
		switch p.InMemory {
		case "redis":
			cacheDockerTemplate = cache_template.DockerComposeRedisConfigTemplate
			cacheVolumeTemplate = cache_template.DockerComposeRedisVolumeTemplate
		case "valkey":
			cacheDockerTemplate = cache_template.DockerComposeValkeyConfigTemplate
			cacheVolumeTemplate = cache_template.DockerComposeValkeyVolumeTemplate
		case "dragonfly":
			cacheDockerTemplate = cache_template.DockerComposeDragonflyConfigTemplate
			cacheVolumeTemplate = cache_template.DockerComposeDragonflyVolumeTemplate
		}

		templates := []string{
			config_template.DockerComposeAppConfigTemplate,
			dbDockerTemplate,
			mqDockerTemplate,
			cacheDockerTemplate,
			objectstorage_template.DockerComposeMinioConfigTemplate,

			// volume
			database_template.DockerComposeDBVolumeTemplate,
			mqVolumetemplate,
			cacheVolumeTemplate,
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
