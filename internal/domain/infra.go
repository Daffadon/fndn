package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	infra_template "github.com/daffadon/fndn/internal/template/infra"
)

func InitQuerierInfra(i infra.CommandRunner, path *string, database *string) error {
	if path != nil {
		folderName := "/internal/infra/storage"
		fileName := folderName + "/querier.go"
		var s string
		switch *database {
		case "postgresql":
			s = infra_template.QuerierPgxInfraTemplate
		case "mariadb":
			s = infra_template.QuerierMariaDBInfraTemplate
		case "clickhouse":
			s = infra_template.QuerierClickHouseInfraTemplate
		case "mongodb", "ferretdb":
			st := struct{ DatabaseName string }{DatabaseName: "database_name"}
			tmp, err := pkg.ParseTemplate(infra_template.QuerierMongoDBInfraTemplate, st)
			if err != nil {
				log.Println(err)
			} else {
				s = tmp
			}
		case "neo4j":
			s = infra_template.QuerierNeo4jInfraTemplate
		}
		if err := pkg.GoFileGenerator(i, path, folderName, fileName, s); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitInMemoryInfra(i infra.CommandRunner, path *string, inMemory *string) error {
	if path != nil {
		folderName := "/internal/infra/cache"
		var fileName, template string
		switch *inMemory {
		case "redis":
			fileName = folderName + "/redis_infra.go"
			template = infra_template.RedisInfraTemplate
		case "valkey":
			fileName = folderName + "/valkey_infra.go"
			template = infra_template.ValkeyInfraTemplate
		case "dragonfly":
			fileName = folderName + "/dragonfly_infra.go"
			template = infra_template.DragonFlyInfraTemplate
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

func InitMQinfra(i infra.CommandRunner, p *Project) error {
	if p.Path != nil {
		folderName := "/internal/infra/mq"
		var fileName, template string
		switch p.MQ {
		case "nats":
			fileName = folderName + "/jetstream_infra.go"
			template = infra_template.JetstreamInfraTemplate
		case "rabbitmq":
			fileName = folderName + "/rabbitmq_infra.go"
			template = infra_template.RabbitMQInfraTemplate
		case "kafka":
			fileName = folderName + "/kafka_infra.go"
			template = infra_template.KafkaMQInfraTemplate
		case "amazon sqs":
			fileName = folderName + "/sqs_infra.go"
			template = infra_template.AmazonSQSInfratemplate
		}
		if fileName != "" || template != "" {
			if err := pkg.GoFileGenerator(i, p.Path, folderName, fileName, template); err != nil {
				log.Fatal(err)
				return err
			}
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
