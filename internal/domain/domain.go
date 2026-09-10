package domain

import (
	"errors"

	"github.com/daffadon/fndn/internal/pkg"
	domain_template "github.com/daffadon/fndn/internal/template/domain"
)

var todoHandlerTemplates = map[string]string{
	"gin":         domain_template.GinTodoHandlerTemplate,
	"chi":         domain_template.ChiTodoHandlerTemplate,
	"echo":        domain_template.EchoTodoHandlerTemplate,
	"fiber":       domain_template.FiberTodoHandlerTemplate,
	"gorilla/mux": domain_template.GorillaTodoHandlerTemplate,
}

var httpHandlerTemplates = map[string]string{
	"gin":         domain_template.GinHTTPHandlerTemplate,
	"chi":         domain_template.ChiHTTPHandlerTemplate,
	"echo":        domain_template.EchoHTTPHandlerTemplate,
	"fiber":       domain_template.FiberHTTPHandlerTemplate,
	"gorilla/mux": domain_template.GorillaHTTPHandlerTemplate,
}

func InitRepositoryDomain(path *string, mn string) error {
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
			return err
		}
		if err := pkg.GoFileGenerator(path, folderName, fileName, c); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitServiceDomain(path *string, mn string) error {
	if path != nil {
		folderName := "/internal/domain/service"
		fileName := folderName + "/todo.go"
		s := struct {
			ModuleName string
		}{
			ModuleName: mn,
		}
		c, err := pkg.ParseTemplate(domain_template.TodoServiceTemplate, s)
		if err != nil {
			return err
		}
		if err := pkg.GoFileGenerator(path, folderName, fileName, c); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitHandlerDomain(path *string, framework *string, mn string) error {
	if path != nil {
		folderName := "/internal/domain/handler"
		fileName := folderName + "/todo.go"
		t, err := lookup(todoHandlerTemplates, "framework", *framework)
		if err != nil {
			return err
		}
		s := struct {
			ModuleName string
		}{
			ModuleName: mn,
		}
		c, err := pkg.ParseTemplate(t, s)
		if err != nil {
			return err
		}
		if err := pkg.GoFileGenerator(path, folderName, fileName, c); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitDTODomain(path *string) error {
	if path != nil {
		folderName := "/internal/domain/dto"
		fileName := folderName + "/todo.go"
		if err := pkg.GoFileGenerator(path, folderName, fileName, domain_template.TodoDTOTemplate); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitHTTPHandlerDomain(path *string, framework *string) error {
	if path != nil {
		folderName := "/internal/domain/handler"
		fileName := folderName + "/http.go"
		t, err := lookup(httpHandlerTemplates, "framework", *framework)
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
