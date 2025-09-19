package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	mq_template "github.com/daffadon/fndn/internal/template/mq"
)

func InitNatsConfig(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := "/config/mq"
		fileName := folderName + "/nats.go"
		if err := pkg.GoFileGenerator(i, path, folderName, fileName, mq_template.NatsConfigTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
