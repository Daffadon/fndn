package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	config_template "github.com/daffadon/fndn/internal/template/config"
)

func InitENVConfig(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := "/config/env"
		fileName := folderName + "/env.go"
		if err := pkg.FileGenerator(i, path, folderName, fileName, config_template.ENVConfigTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
