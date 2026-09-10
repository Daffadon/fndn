package app

import (
	"fmt"

	"github.com/daffadon/fndn/internal/domain"
	"github.com/daffadon/fndn/internal/infra"
)

type InitGenerate struct {
	Runner infra.CommandRunner
}

func (i *InitGenerate) generators() map[domain.GeneratorType]func(value, path string) error {
	return map[domain.GeneratorType]func(string, string) error{
		domain.GeneratorFramework: domain.GenerateSpecificFramework,
		domain.GeneratorDatabase:  domain.GenerateSpecificDatabase,
		domain.GeneratorMQ:        domain.GenerateSpecificMQ,
		domain.GeneratorCache:     domain.GenerateSpecificCache,
		domain.GeneratorStorage:   domain.GenerateSpecificStorage,
	}
}

func (i *InitGenerate) Run(g *domain.Generator, progressCh chan<- string) error {
	path := "."
	gen, ok := i.generators()[g.Type]
	if !ok {
		return fmt.Errorf("unknown generator type %q", g.Type)
	}
	progressCh <- fmt.Sprintf("Running %s generation", g.Type)
	if err := gen(g.Value, path); err != nil {
		return err
	}
	return i.Finalize(path, progressCh)
}

// Finalize runs toolchain commands needing network. Split from Run's
// file generation so scaffold stays testable offline.
func (i *InitGenerate) Finalize(path string, progressCh chan<- string) error {
	progressCh <- "Running go get ./... to download 3rd party modules"
	if err := i.Runner.Run("go", []string{"get", "./..."}, path); err != nil {
		return err
	}
	progressCh <- "Running go mod tidy"
	if err := i.Runner.Run("go", []string{"mod", "tidy"}, path); err != nil {
		return err
	}
	return nil
}
