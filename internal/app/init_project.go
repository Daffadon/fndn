package app

import (
	"sync"

	"github.com/daffadon/fndn/internal/domain"
	"github.com/daffadon/fndn/internal/infra"
)

type InitProjectUseCase struct {
	Runner infra.CommandRunner
}

func (uc *InitProjectUseCase) Run(p *domain.Project) error {
	if p.Path == "" {
		newPath := p.Name
		if err := uc.Runner.Run("mkdir", []string{newPath}, ""); err != nil {
			return err
		}
		p.Path = newPath
	}
	// create project and init
	if err := domain.InitProject(uc.Runner, p.Path, p.ModuleName); err != nil {
		return err
	}
	// init git or not
	if err := domain.InitGit(uc.Runner, &p.Path, p.Git); err != nil {
		return err
	}
	// run each init in a goroutine
	initFuncs := []func() error{
		func() error { return domain.InitGin(uc.Runner, &p.Path) },
		func() error { return domain.InitENVConfig(uc.Runner, &p.Path) },
		func() error { return domain.InitZerologConfig(uc.Runner, &p.Path) },
		func() error { return domain.InitRedisConfig(uc.Runner, &p.Path) },
		func() error { return domain.InitNatsConfig(uc.Runner, &p.Path) },
		func() error { return domain.InitMinioConfig(uc.Runner, &p.Path) },
		func() error { return domain.InitPostgresqlConfig(uc.Runner, &p.Path) },
		func() error { return domain.InitDependencyInjection(uc.Runner, &p.Path) },
		func() error { return domain.InitBootStrap(uc.Runner, &p.Path) },
		func() error { return domain.InitServer(uc.Runner, &p.Path) },
		func() error { return domain.InitMain(uc.Runner, &p.Path) },
		func() error { return domain.InitAirConfig(uc.Runner, &p.Path, p.Air) },
	}

	errCh := make(chan error, len(initFuncs))
	var wg sync.WaitGroup

	for _, f := range initFuncs {
		wg.Add(1)
		go func(fn func() error) {
			defer wg.Done()
			if err := fn(); err != nil {
				errCh <- err
			}
		}(f)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
		}
	}
	if err := uc.Runner.Run("goimports", []string{"-w", "."}, p.Path); err != nil {
		return err
	}
	if err := uc.Runner.Run("go", []string{"mod", "tidy"}, p.Path); err != nil {
		return err
	}

	return nil
}
