package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	framework_template "github.com/daffadon/fndn/internal/template/framework"
)

func InitGin(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := "/config/router"
		fileName := folderName + "/http.go"
		if err := pkg.FileGenerator(i, path, folderName, fileName, framework_template.GinConfigTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
