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
		fileName := folderName + "/mq.go"
		var template string
		switch p.MQ {
		case "nats":
			template = mq_template.NatsConfigTemplate
		case "rabbitmq":
			template = mq_template.RabbitMQConfigTemplate
		case "kafka":
			template = mq_template.KafkaConfigTemplate
		case "amazon sqs":
			template = mq_template.AmazonSQSConfigTemplate
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

func GenerateSpecificMQ(mq string, infraRunner infra.CommandRunner, path string) error {
	// check folder config/router/ exist or not
	// check filename
	folderName := "/config/mq"
	fileName := folderName + "/mq.go"

	// if exist, the file name add _framework_name
	exist := pkg.IsFileExists("." + fileName)
	if exist {
		fileName = folderName + "/mq_" + mq + ".go"
	}

	var t string
	switch mq {
	case "nats":
		t = mq_template.NatsConfigTemplate
	case "rabbitmq":
		t = mq_template.RabbitMQConfigTemplate
	case "kafka":
		t = mq_template.KafkaConfigTemplate
	case "amazon sqs":
		t = mq_template.AmazonSQSConfigTemplate
	}

	if err := pkg.GoFileGenerator(infraRunner, &path, folderName, fileName, t); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
