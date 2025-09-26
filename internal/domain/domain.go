package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	domain_template "github.com/daffadon/fndn/internal/template/domain"
)

func InitRepositoryDomain(i infra.CommandRunner, path *string, mn string) error {
	if path != nil {
		folderName := "/internal/domain/repository"
		fileName := folderName + "/todo.go"
		s := struct {
			ModuleName string
		}{
			ModuleName: mn,
		}
		c, err := pkg.ParseTemplate(domain_template.TodoRepositoryTemplate, s)
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

func InitServiceDomain(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := "/internal/domain/service"
		fileName := folderName + "/todo.go"
		if err := pkg.GoFileGenerator(i, path, folderName, fileName, domain_template.TodoServiceTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitHandlerDomain(i infra.CommandRunner, path *string, framework *string) error {
	if path != nil {
		folderName := "/internal/domain/handler"
		fileName := folderName + "/todo.go"
		var t string
		switch *framework {
		case "gin":
			t = domain_template.GinTodoHandlerTemplate
		case "chi":
			t = domain_template.ChiTodoHandlerTemplate
		case "echo":
			t = domain_template.EchoTodoHandlerTemplate
		case "fiber":
			t = domain_template.FiberTodoHandlerTemplate
		case "gorilla/mux":
			t = domain_template.GorillaTodoHandlerTemplate
		}
		if err := pkg.GoFileGenerator(i, path, folderName, fileName, t); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitDTODomain(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := "/internal/domain/dto"
		fileName := folderName + "/todo.go"
		if err := pkg.GoFileGenerator(i, path, folderName, fileName, domain_template.TodoDTOTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitHTTPHandlerDomain(i infra.CommandRunner, path *string, framework *string) error {
	if path != nil {
		folderName := "/internal/domain/handler"
		fileName := folderName + "/http.go"
		var t string
		switch *framework {
		case "gin":
			t = domain_template.GinHTTPHandlerTemplate
		case "chi":
			t = domain_template.ChiHTTPHandlerTemplate
		case "echo":
			t = domain_template.EchoHTTPHandlerTemplate
		case "fiber":
			t = domain_template.FiberHTTPHandlerTemplate
		case "gorilla/mux":
			t = domain_template.GorillaHTTPHandlerTemplate
		}
		if err := pkg.GoFileGenerator(i, path, folderName, fileName, t); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
