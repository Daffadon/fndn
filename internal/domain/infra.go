package domain

import (
	"errors"

	"github.com/daffadon/fndn/internal/pkg"
	infra_template "github.com/daffadon/fndn/internal/template/infra"
)

var querierTemplates = map[string]string{
	"postgresql": infra_template.QuerierPgxInfraTemplate,
	"mariadb":    infra_template.QuerierMariaDBInfraTemplate,
	"clickhouse": infra_template.QuerierClickHouseInfraTemplate,
	"neo4j":      infra_template.QuerierNeo4jInfraTemplate,
}

var inMemoryInfraFiles = map[string][2]string{
	"redis":     {"/redis_infra.go", infra_template.RedisInfraTemplate},
	"valkey":    {"/valkey_infra.go", infra_template.ValkeyInfraTemplate},
	"dragonfly": {"/dragonfly_infra.go", infra_template.DragonFlyInfraTemplate},
	"redict":    {"/redict_infra.go", infra_template.RedictInfraTemplate},
}

var mqInfraFiles = map[string][2]string{
	"nats":       {"/jetstream_infra.go", infra_template.JetstreamInfraTemplate},
	"rabbitmq":   {"/rabbitmq_infra.go", infra_template.RabbitMQInfraTemplate},
	"kafka":      {"/kafka_infra.go", infra_template.KafkaMQInfraTemplate},
	"amazon sqs": {"/sqs_infra.go", infra_template.AmazonSQSInfratemplate},
}

var objectStorageInfraFiles = map[string][2]string{
	"rustfs":    {"/rustfs.go", infra_template.RustfsInfraTemplate},
	"seaweedfs": {"/seaweedfs.go", infra_template.SeaweedInfraTemplate},
	"minio":     {"/minio.go", infra_template.MinioInfraTemplate},
}

func InitQuerierInfra(path *string, database *string) error {
	if *database == None {
		return nil
	}
	if path != nil {
		folderName := "/internal/infra/storage"
		fileName := folderName + "/querier.go"
		var s string
		switch *database {
		case "mongodb", "ferretdb":
			st := struct{ DatabaseName string }{DatabaseName: "database_name"}
			tmp, err := pkg.ParseTemplate(infra_template.QuerierMongoDBInfraTemplate, st)
			if err != nil {
				return err
			}
			s = tmp
		default:
			var err error
			s, err = lookup(querierTemplates, "database", *database)
			if err != nil {
				return err
			}
		}
		if err := pkg.GoFileGenerator(path, folderName, fileName, s); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitInMemoryInfra(path *string, inMemory *string) error {
	if *inMemory == None {
		return nil
	}
	if path != nil {
		folderName := "/internal/infra/cache"
		file, err := lookupFile(inMemoryInfraFiles, "in-memory store", *inMemory)
		if err != nil {
			return err
		}
		if err := pkg.GoFileGenerator(path, folderName, folderName+file[0], file[1]); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitMQinfra(p *Project) error {
	if p.MQ == None {
		return nil
	}
	if p.Path != nil {
		folderName := "/internal/infra/mq"
		file, err := lookupFile(mqInfraFiles, "message queue", p.MQ)
		if err != nil {
			return err
		}
		if err := pkg.GoFileGenerator(p.Path, folderName, folderName+file[0], file[1]); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitObjectStorageInfra(path, os *string) error {
	if *os == None {
		return nil
	}
	if path != nil {
		folderName := "/internal/infra/storage"
		file, err := lookupFile(objectStorageInfraFiles, "object storage", *os)
		if err != nil {
			return err
		}
		if err := pkg.GoFileGenerator(path, folderName, folderName+file[0], file[1]); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitMinioInfra(path *string) error {
	if path != nil {
		folderName := "/internal/infra/storage"
		fileName := folderName + "/minio.go"
		if err := pkg.GoFileGenerator(path, folderName, fileName, infra_template.MinioInfraTemplate); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
