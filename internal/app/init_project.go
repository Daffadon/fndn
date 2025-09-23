package app

import (
	"os"
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
		if err := os.MkdirAll(newPath, 0755); err != nil {
			return err
		}
		p.Path = newPath
	}
	// init
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
		// config
		func() error { return domain.InitGin(uc.Runner, &p.Path) },
		func() error { return domain.InitENVConfig(uc.Runner, &p.Path) },
		func() error { return domain.InitZerologConfig(uc.Runner, &p.Path) },
		func() error { return domain.InitRedisConfig(uc.Runner, &p.Path) },
		func() error { return domain.InitNatsConfig(uc.Runner, &p.Path) },
		func() error { return domain.InitMinioConfig(uc.Runner, &p.Path) },
		func() error { return domain.InitPostgresqlConfig(uc.Runner, &p.Path) },

		// infra
		func() error { return domain.InitQuerierInfra(uc.Runner, &p.Path) },
		func() error { return domain.InitRedisInfra(uc.Runner, &p.Path) },
		func() error { return domain.InitJetstreamInfra(uc.Runner, &p.Path) },
		func() error { return domain.InitMinioInfra(uc.Runner, &p.Path) },

		// domain
		func() error { return domain.InitDTODomain(uc.Runner, &p.Path) },
		func() error { return domain.InitRepositoryDomain(uc.Runner, &p.Path, p.ModuleName) },
		func() error { return domain.InitServiceDomain(uc.Runner, &p.Path) },
		func() error { return domain.InitHandlerDomain(uc.Runner, &p.Path) },
		func() error { return domain.InitHTTPHandlerDomain(uc.Runner, &p.Path) },

		// cmd
		func() error { return domain.InitDependencyInjection(uc.Runner, &p.Path) },
		func() error { return domain.InitBootStrap(uc.Runner, &p.Path) },
		func() error { return domain.InitServer(uc.Runner, &p.Path) },
		func() error { return domain.InitMain(uc.Runner, &p.Path) },

		// global config
		func() error { return domain.InitAirConfig(uc.Runner, &p.Path, p.Air) },
		func() error { return domain.InitYamlConfig(uc.Runner, p, &p.Path) },
		func() error { return domain.InitGitignoreConfig(uc.Runner, &p.Path) },
		func() error { return domain.InitDockerFileConfig(uc.Runner, &p.Path, p.Name) },
		func() error { return domain.InitDockerComposeConfig(uc.Runner, &p.Path, p.Name) },
		func() error { return domain.InitNatsConfigFile(uc.Runner, &p.Path) },
		func() error { return domain.InitDotEnvExampleConfig(uc.Runner, &p.Path) },
		func() error { return domain.InitReadme(uc.Runner, &p.Path) },
		func() error { return domain.InitVersion(uc.Runner, &p.Path) },
		func() error { return domain.InitBuildScript(uc.Runner, &p.Path, p.ModuleName) },
		func() error { return domain.InitBinaryBuildScript(uc.Runner, &p.Path, p.Name) },
		func() error { return domain.InitMakefile(uc.Runner, &p.Path) },
	}

	errCh := make(chan error, len(initFuncs))
	var wg sync.WaitGroup
	sem := make(chan struct{}, 5)

	for _, f := range initFuncs {
		wg.Add(1)
		sem <- struct{}{}
		go func(fn func() error) {
			defer wg.Done()
			defer func() { <-sem }()
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
	if err := uc.Runner.Run("go", []string{"get", "-u", "./..."}, p.Path); err != nil {
		return err
	}
	if err := uc.Runner.Run("go", []string{"mod", "tidy"}, p.Path); err != nil {
		return err
	}

	return nil
}
