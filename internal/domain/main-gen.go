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
			DBConnection string
			MQInit       string
			MQInfra      string
		}
		switch p.Database {
		case "postgresql", "mariadb", "clickhouse":
			st.DBConnection = "NewSQLConn"
		case "neo4j":
			st.DBConnection = "NewGraphDBConn"
		default:
			st.DBConnection = "NewNoSQLConn"
		}

		provideMQDefault := `
			// mq client connection
			if err := container.Provide(mq.NewMQConnection); err != nil {
				panic("Failed to provide mq connection: " + err.Error())
			}
		`
		switch p.MQ {
		case "nats":
			st.MQInfra = "NewJetstreamInfra"
			provideMQDefault += `
				// jetstream connection
				if err := container.Provide(jetstream.New); err != nil {
					panic("Failed to provide jetstream instance: " + err.Error())
				}
			`
		case "rabbitmq":
			st.MQInfra = "NewRabbitMQInfra"
		case "kafka":
			st.MQInfra = "NewKafkaInfra"
		case "amazon sqs":
			st.MQInfra = "NewSQSInfra"
		}
		st.MQInit = provideMQDefault

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
		c, err := pkg.HTTPServerParser(p.Framework, p.Database, p.MQ)
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
