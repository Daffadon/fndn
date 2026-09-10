package domain

import (
	"errors"

	"github.com/daffadon/fndn/internal/pkg"
	main_template "github.com/daffadon/fndn/internal/template/main"
)

var diFrameworkInits = map[string]string{
	"gin":         "NewHTTPGin",
	"chi":         "NewHTTPChi",
	"echo":        "NewHTTPEcho",
	"fiber":       "NewHTTPFiber",
	"gorilla/mux": "NewHTTPMux",
}

var diDBConns = map[string]string{
	"postgresql": "NewPostgresqlConn",
	"mariadb":    "NewMariaDBConn",
	"clickhouse": "NewClickhouseConn",
	"mongodb":    "NewMongoDBConn",
	"ferretdb":   "NewFerretDBConn",
	"neo4j":      "NewNeoFourJConn",
}

// diMQ: init snippet + infra constructor per message queue.
var diMQs = map[string][2]string{
	"nats": {"NewJetstreamInfra", `
			// mq client connection
			if err := container.Provide(mq.NewNatsConnection); err != nil {
				panic("Failed to provide mq connection: " + err.Error())
			}
			// jetstream connection
			if err := container.Provide(jetstream.New); err != nil {
				panic("Failed to provide jetstream instance: " + err.Error())
			}
		`},
	"rabbitmq": {"NewRabbitMQInfra", `
			// mq client connection
			if err := container.Provide(mq.NewRabbitMQConnection); err != nil {
				panic("Failed to provide mq connection: " + err.Error())
			}
		`},
	"kafka": {"NewKafkaInfra", `
			// mq client connection
			if err := container.Provide(mq.NewKafkaConnection); err != nil {
				panic("Failed to provide mq connection: " + err.Error())
			}
		`},
	"amazon sqs": {"NewSQSInfra", `
			// mq client connection
			if err := container.Provide(mq.NewSQSConnection); err != nil {
				panic("Failed to provide mq connection: " + err.Error())
			}
		`},
}

var diCacheConns = map[string][2]string{
	"redis":     {"NewRedisConnection", "NewRedisCache"},
	"valkey":    {"NewValkeyConnection", "NewValkeyCache"},
	"dragonfly": {"NewDragonflyConnection", "NewDragonflyCache"},
	"redict":    {"NewRedictConnection", "NewRedictCache"},
}

var diOSConns = map[string][2]string{
	"rustfs":    {"NewRustfsConnection", "NewRustfsInfra"},
	"seaweedfs": {"NewSeaweedfsConnection", "NewSeaweedfsInfra"},
	"minio":     {"NewMinioConnection", "NewMinioInfra"},
}

func InitDependencyInjection(p *Project) error {
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
			HasDB           bool
			HasMQ           bool
			HasCache        bool
			HasOS           bool
		}
		var err error
		if st.HTTPInit, err = lookup(diFrameworkInits, "framework", p.Framework); err != nil {
			return err
		}
		if st.DBConnection, err = section(diDBConns, "database", p.Database); err != nil {
			return err
		}
		st.HasDB = p.Database != None
		mq, err := sectionFile(diMQs, "message queue", p.MQ)
		if err != nil {
			return err
		}
		st.MQInfra, st.MQInit = mq[0], mq[1]
		st.HasMQ = p.MQ != None
		cache, err := sectionFile(diCacheConns, "in-memory store", p.InMemory)
		if err != nil {
			return err
		}
		st.CacheConnection, st.CacheInfra = cache[0], cache[1]
		st.HasCache = p.InMemory != None
		os, err := sectionFile(diOSConns, "object storage", p.ObjectStorage)
		if err != nil {
			return err
		}
		st.OSConnection, st.OSInfra = os[0], os[1]
		st.HasOS = p.ObjectStorage != None

		template, err := pkg.ParseTemplate(main_template.DITemplate, st)
		if err != nil {
			return err
		}
		if err := pkg.GoFileGenerator(p.Path, folderName, fileName, template); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitBootStrap(path *string) error {
	if path != nil {
		folderName := "/cmd/bootstrap"
		fileName := folderName + "/bootstrap.go"
		if err := pkg.GoFileGenerator(path, folderName, fileName, main_template.BootStrapTemplate); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitServer(p *Project) error {
	if p.Path != nil {
		folderName := "/cmd/server"
		fileName := folderName + "/server.go"
		c, err := pkg.HTTPServerParser(p.Framework, p.Database, p.MQ, p.InMemory)
		if err != nil {
			return err
		}
		if err := pkg.GoFileGenerator(p.Path, folderName, fileName, c); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitMain(path *string) error {
	if path != nil {
		folderName := "/cmd"
		fileName := folderName + "/main.go"
		if err := pkg.GoFileGenerator(path, folderName, fileName, main_template.MainTemplate); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
