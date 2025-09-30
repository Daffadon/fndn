package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	main_template "github.com/daffadon/fndn/internal/template/main"
)

func InitDependencyInjection(i infra.CommandRunner, p *Project) error {
	if p.Path != nil {
		folderName := "/cmd/di"
		fileName := folderName + "/container.go"
		var st struct {
			DBConnection string
		}
		switch p.Database {
		case "postgresql", "mariadb":
			st.DBConnection = "NewSQLConn"
		default:
			st.DBConnection = "NewNoSQLConn"
		}
		template, err := pkg.ParseTemplate(main_template.DITemplate, st)
		if err != nil {
			log.Fatal(err)
			return err
		}
		if err := pkg.GoFileGenerator(i, p.Path, folderName, fileName, template); err != nil {
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

func InitServer(i infra.CommandRunner, d *Project) error {
	if d.Path != nil {
		folderName := "/cmd/server"
		fileName := folderName + "/server.go"
		c, err := pkg.HTTPServerParser(d.Framework, d.Database)
		if err != nil {
			log.Fatal(err)
			return err
		}
		if err := pkg.GoFileGenerator(i, d.Path, folderName, fileName, c); err != nil {
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
