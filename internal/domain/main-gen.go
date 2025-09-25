package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	main_template "github.com/daffadon/fndn/internal/template/main"
)

func InitDependencyInjection(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := "/cmd/di"
		fileName := folderName + "/container.go"
		if err := pkg.GoFileGenerator(i, path, folderName, fileName, main_template.DITemplate); err != nil {
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

func InitServer(i infra.CommandRunner, path *string, fwk *string) error {
	if path != nil {
		folderName := "/cmd/server"
		fileName := folderName + "/server.go"
		c, err := pkg.HTTPServerParser(*fwk)
		if err != nil {
			log.Fatal(err)
			return err
		}
		if err := pkg.GoFileGenerator(i, path, folderName, fileName, c); err != nil {
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
