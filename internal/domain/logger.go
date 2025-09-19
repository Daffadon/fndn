package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	logger_template "github.com/daffadon/fndn/internal/template/logger"
)

func InitZerologConfig(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := "/config/logger"
		fileName := folderName + "/zerolog.go"
		if err := pkg.GoFileGenerator(i, path, folderName, fileName, logger_template.ZerologConfigTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
