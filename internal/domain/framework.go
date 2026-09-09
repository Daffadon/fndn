package domain

import (
	"errors"

	"github.com/daffadon/fndn/internal/pkg"
	framework_template "github.com/daffadon/fndn/internal/template/framework"
)

var frameworkTemplates = map[string]string{
	"gin":         framework_template.GinConfigTemplate,
	"chi":         framework_template.ChiConfigTemplate,
	"echo":        framework_template.EchoConfigTemplate,
	"fiber":       framework_template.FiberConfigTemplate,
	"gorilla/mux": framework_template.GorillaMuxConfigTemplate,
}

func InitFramework(path *string, framework *string) error {
	if path != nil {
		folderName := "/config/router"
		fileName := folderName + "/http.go"
		t, err := lookup(frameworkTemplates, "framework", *framework)
		if err != nil {
			return err
		}
		if err := pkg.GoFileGenerator(path, folderName, fileName, t); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func GenerateSpecificFramework(framework string, path string) error {
	folderName := "/config/router"
	fileName := resolveTarget(folderName+"/http.go", folderName+"/http", framework)

	t, err := lookup(frameworkTemplates, "framework", framework)
	if err != nil {
		return err
	}
	if err := pkg.GoFileGenerator(&path, folderName, fileName, t); err != nil {
		return err
	}
	return nil
}
