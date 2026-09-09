package domain

import (
	"errors"

	"github.com/daffadon/fndn/internal/pkg"
	mq_template "github.com/daffadon/fndn/internal/template/mq"
)

var mqTemplates = map[string]string{
	"nats":       mq_template.NatsConfigTemplate,
	"rabbitmq":   mq_template.RabbitMQConfigTemplate,
	"kafka":      mq_template.KafkaConfigTemplate,
	"amazon sqs": mq_template.AmazonSQSConfigTemplate,
}

var mqConfigFiles = map[string][2]string{
	"nats":     {"/nats-server.conf", mq_template.NatsConfigFileTemplate},
	"rabbitmq": {"/definition.json", mq_template.RabbitMQConfigFileTemplate},
	"kafka":    {"/jaas.conf", mq_template.KafkaConfigFileTemplate},
}

func InitMQConfig(p *Project) error {
	if p.Path != nil {
		folderName := "/config/mq"
		fileName := folderName + "/mq.go"
		template, err := lookup(mqTemplates, "message queue", p.MQ)
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

func InitMQConfigFile(p *Project) error {
	if p.Path != nil {
		folderName := "/config/mq"
		file, ok := mqConfigFiles[p.MQ]
		if !ok {
			return nil
		}
		if err := pkg.GenericFileGenerator(p.Path, folderName, folderName+file[0], file[1]); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func GenerateSpecificMQ(mq string, path string) error {
	folderName := "/config/mq"
	fileName := resolveTarget(folderName+"/mq.go", folderName+"/mq", mq)

	t, err := lookup(mqTemplates, "message queue", mq)
	if err != nil {
		return err
	}
	if err := pkg.GoFileGenerator(&path, folderName, fileName, t); err != nil {
		return err
	}
	return nil
}
