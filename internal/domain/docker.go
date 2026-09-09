package domain

import (
	"errors"

	"github.com/daffadon/fndn/internal/pkg"
	cache_template "github.com/daffadon/fndn/internal/template/cache"
	config_template "github.com/daffadon/fndn/internal/template/config"
	database_template "github.com/daffadon/fndn/internal/template/database"
	mq_template "github.com/daffadon/fndn/internal/template/mq"
	objectstorage_template "github.com/daffadon/fndn/internal/template/object_storage"
)

var dockerDBTemplates = map[string]string{
	"postgresql": database_template.DockerComposePostgresqlConfigTemplate,
	"mariadb":    database_template.DockerComposeMariaDBConfigTemplate,
	"clickhouse": database_template.DockerComposeClickHouseConfigTemplate,
	"mongodb":    database_template.DockerComposeMongoDBConfigTemplate,
	"ferretdb":   database_template.DockerComposeFerretDBConfigTemplate,
	"neo4j":      database_template.DockerComposeNeo4JConfigTemplate,
}

var dockerMQTemplates = map[string][2]string{
	"nats":     {mq_template.DockerComposeNatsConfigTemplate, mq_template.DockerComposeNatsVolumeTemplate},
	"rabbitmq": {mq_template.DockerComposeRabbitMQConfigTemplate, mq_template.DockerComposeRabbitVolumeTemplate},
	"kafka":    {mq_template.DockerComposeKafkaConfigTemplate, mq_template.DockerComposeKafkaVolumeTemplate},
}

var dockerCacheTemplates = map[string][2]string{
	"redis":     {cache_template.DockerComposeRedisConfigTemplate, cache_template.DockerComposeRedisVolumeTemplate},
	"valkey":    {cache_template.DockerComposeValkeyConfigTemplate, cache_template.DockerComposeValkeyVolumeTemplate},
	"dragonfly": {cache_template.DockerComposeDragonflyConfigTemplate, cache_template.DockerComposeDragonflyVolumeTemplate},
	"redict":    {cache_template.DockerComposeRedictConfigTemplate, cache_template.DockerComposeRedictVolumeTemplate},
}

var dockerOSTemplates = map[string][2]string{
	"rustfs":    {objectstorage_template.DockerComposeRustfsConfigTemplate, objectstorage_template.DockerComposeRustfsVolumeTemplate},
	"seaweedfs": {objectstorage_template.DockerComposeSeaweedfsConfigTemplate, objectstorage_template.DockerComposeSeaweedfsVolumeTemplate},
	"minio":     {objectstorage_template.DockerComposeMinioConfigTemplate, objectstorage_template.DockerComposeMinioVolumeTemplate},
}

func InitDockerFileConfig(path *string, projectName string) error {
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
			return err
		}
		if err := pkg.GenericFileGenerator(path, folderName, fileName, c); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitDockerComposeConfig(p *Project) error {
	if p.Path != nil {
		folderName := ""
		fileName := folderName + "/docker-compose.yml"
		st := struct {
			ProjectName string
		}{
			ProjectName: p.Name,
		}
		var results []string
		dbDockerTemplate, err := lookup(dockerDBTemplates, "database", p.Database)
		if err != nil {
			return err
		}
		mqDocker, err := lookupFile(dockerMQTemplates, "message queue", p.MQ)
		if err != nil {
			return err
		}
		cacheDocker, err := lookupFile(dockerCacheTemplates, "in-memory store", p.InMemory)
		if err != nil {
			return err
		}
		osDocker, err := lookupFile(dockerOSTemplates, "object storage", p.ObjectStorage)
		if err != nil {
			return err
		}

		templates := []string{
			config_template.DockerComposeAppConfigTemplate,
			dbDockerTemplate,
			mqDocker[0],
			cacheDocker[0],
			osDocker[0],

			// volume
			database_template.DockerComposeDBVolumeTemplate,
			mqDocker[1],
			cacheDocker[1],
			osDocker[1],
		}
		for _, tpl := range templates {
			c, err := pkg.ParseTemplate(tpl, st)
			if err != nil {
				return err
			}
			results = append(results, c)
		}

		s := config_template.DockerComposeDefaultConfigTemplate
		for i := range results {
			if i == 5 {
				s += config_template.DockerComposeVolumeConfigTemplate
			}
			s += results[i]
		}
		if err := pkg.GenericFileGenerator(p.Path, folderName, fileName, s); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
