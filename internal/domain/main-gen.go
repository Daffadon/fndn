package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	main_template "github.com/daffadon/fndn/internal/template/main"
)

func InitDependencyInjection(i infra.CommandRunner, p *Project) error {
	if p.Path != nil {
		folderName := "/cmd/di"
		fileName := folderName + "/container.go"
		var st struct {
			HTTPInit        string
			DBConnection    string
			MQInit          string
			MQInfra         string
			CacheConnection string
			CacheInfra      string
			OSConnection    string
			OSInfra         string
		}
		switch p.Framework {
		case "gin":
			st.HTTPInit = "NewHTTPGin"
		case "chi":
			st.HTTPInit = "NewHTTPChi"
		case "echo":
			st.HTTPInit = "NewHTTPEcho"
		case "fiber":
			st.HTTPInit = "NewHTTPFiber"
		case "gorilla/mux":
			st.HTTPInit = "NewHTTPMux"
		}
		switch p.Database {
		case "postgresql":
			st.DBConnection = "NewPostgresqlConn"
		case "mariadb":
			st.DBConnection = "NewMariaDBConn"
		case "clickhouse":
			st.DBConnection = "NewClickhouseConn"
		case "mongodb":
			st.DBConnection = "NewMongoDBConn"
		case "ferretdb":
			st.DBConnection = "NewFerretDBConn"
		case "neo4j":
			st.DBConnection = "NewNeoFourJConn"
		}

		switch p.MQ {
		case "nats":
			st.MQInfra = "NewJetstreamInfra"
			st.MQInit = `
				// mq client connection
				if err := container.Provide(mq.NewNatsConnection); err != nil {
					panic("Failed to provide mq connection: " + err.Error())
				}
				// jetstream connection
				if err := container.Provide(jetstream.New); err != nil {
					panic("Failed to provide jetstream instance: " + err.Error())
				}
			`
		case "rabbitmq":
			st.MQInit = `
				// mq client connection
				if err := container.Provide(mq.NewRabbitMQConnection); err != nil {
					panic("Failed to provide mq connection: " + err.Error())
				}
			`
			st.MQInfra = "NewRabbitMQInfra"
		case "kafka":
			st.MQInit = `
				// mq client connection
				if err := container.Provide(mq.NewKafkaConnection); err != nil {
					panic("Failed to provide mq connection: " + err.Error())
				}
			`
			st.MQInfra = "NewKafkaInfra"
		case "amazon sqs":
			st.MQInit = `
				// mq client connection
				if err := container.Provide(mq.NewSQSConnection); err != nil {
					panic("Failed to provide mq connection: " + err.Error())
				}
			`
			st.MQInfra = "NewSQSInfra"
		}

		switch p.InMemory {
		case "redis":
			st.CacheConnection = "NewRedisConnection"
			st.CacheInfra = "NewRedisCache"
		case "valkey":
			st.CacheConnection = "NewValkeyConnection"
			st.CacheInfra = "NewValkeyCache"
		case "dragonfly":
			st.CacheConnection = "NewDragonflyConnection"
			st.CacheInfra = "NewDragonflyCache"
		case "redict":
			st.CacheConnection = "NewRedictConnection"
			st.CacheInfra = "NewRedictCache"
		}
		switch p.ObjectStorage {
		case "rustfs":
			st.OSConnection = "NewRustfsConnection"
			st.OSInfra = "NewRustfsInfra"

		case "seaweedfs":
			st.OSConnection = "NewSeaweedfsConnection"
			st.OSInfra = "NewSeaweedfsInfra"

		case "minio":
			st.OSConnection = "NewMinioConnection"
			st.OSInfra = "NewMinioInfra"
		}

		template, err := pkg.ParseTemplate(main_template.DITemplate, st)
		if err != nil {
			log.Fatal(err)
			return err
		}
		if err := pkg.GoFileGenerator(i, p.Path, folderName, fileName, template); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitBootStrap(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := "/cmd/bootstrap"
		fileName := folderName + "/bootstrap.go"
		if err := pkg.GoFileGenerator(i, path, folderName, fileName, main_template.BootStrapTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitServer(i infra.CommandRunner, p *Project) error {
	if p.Path != nil {
		folderName := "/cmd/server"
		fileName := folderName + "/server.go"
		c, err := pkg.HTTPServerParser(p.Framework, p.Database, p.MQ, p.InMemory)
		if err != nil {
			log.Fatal(err)
			return err
		}
		if err := pkg.GoFileGenerator(i, p.Path, folderName, fileName, c); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitMain(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := "/cmd"
		fileName := folderName + "/main.go"
		if err := pkg.GoFileGenerator(i, path, folderName, fileName, main_template.MainTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
