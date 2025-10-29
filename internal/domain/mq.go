package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	mq_template "github.com/daffadon/fndn/internal/template/mq"
)

func InitMQConfig(i infra.CommandRunner, p *Project) error {
	if p.Path != nil {
		folderName := "/config/mq"
		var fileName string
		var template string
		switch p.MQ {
		case "nats":
			fileName = folderName + "/nats.go"
			template = mq_template.NatsConfigTemplate
		case "rabbitmq":
			fileName = folderName + "/rabbitmq.go"
			template = mq_template.RabbitMQConfigTemplate
		case "kafka":
			fileName = folderName + "/kafka.go"
			template = mq_template.KafkaConfigTemplate
		case "amazon sqs":
			fileName = folderName + "/sqs.go"
			template = mq_template.AmazonSQSConfigTemplate
		}
		if err := pkg.GoFileGenerator(i, p.Path, folderName, fileName, template); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitMQConfigFile(i infra.CommandRunner, p *Project) error {
	if p.Path != nil {
		folderName := "/config/mq"
		var fileName, template string
		switch p.MQ {
		case "nats":
			fileName = folderName + "/nats-server.conf"
			template = mq_template.NatsConfigFileTemplate
		case "rabbitmq":
			fileName = folderName + "/definition.json"
			template = mq_template.RabbitMQConfigFileTemplate
		case "kafka":
			fileName = folderName + "/jaas.conf"
			template = mq_template.KafkaConfigFileTemplate
		}
		if fileName != "" || template != "" {
			if err := pkg.GenericFileGenerator(i, p.Path, folderName, fileName, template); err != nil {
				log.Fatal(err)
				return err
			}
		}
		return nil
	}
	return errors.New("path is nil")
}
