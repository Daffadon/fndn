package app

import (
	"github.com/daffadon/fndn/internal/domain"
	"github.com/daffadon/fndn/internal/infra"
)

type InitGenerate struct {
	Runner infra.CommandRunner
}

func (i *InitGenerate) Run(g *domain.Generator, progressCh chan<- string) error {
	path := "."
	switch g.Type {
	case "framework":
		progressCh <- "Running framework generation"
		if err := domain.GenerateSpecificFramework(g.Value, i.Runner, path); err != nil {
			return err
		}
	case "database":
		progressCh <- "Running database config generation"
		if err := domain.GenerateSpecificDatabase(g.Value, i.Runner, path); err != nil {
			return err
		}
	}
	progressCh <- "Running go get -u ./... to download 3rd party modules"
	if err := i.Runner.Run("go", []string{"get", "-u", "./..."}, path); err != nil {
		return err
	}
	progressCh <- "Running go mod tidy"
	if err := i.Runner.Run("go", []string{"mod", "tidy"}, path); err != nil {
		return err
	}
	return nil
}
