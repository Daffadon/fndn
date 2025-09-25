package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	framework_template "github.com/daffadon/fndn/internal/template/framework"
)

func InitFramework(i infra.CommandRunner, path *string, framework *string) error {
	if path != nil {
		folderName := "/config/router"
		fileName := folderName + "/http.go"
		var t string
		switch *framework {
		case "gin":
			t = framework_template.GinConfigTemplate
		case "chi":
			t = framework_template.ChiConfigTemplate
		case "echo":
			t = framework_template.EchoConfigTemplate
		}
		if err := pkg.GoFileGenerator(i, path, folderName, fileName, t); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
