package app

import (
	"os"

	"github.com/daffadon/fndn/internal/domain"
	"github.com/daffadon/fndn/internal/infra"
)

type InitProjectUseCase struct {
	Runner infra.CommandRunner
}

type initStep struct {
	progress string
	run      func() error
}

func (uc *InitProjectUseCase) Run(p *domain.Project, progressCh chan<- string) error {
	if *p.Path == "" {
		newPath := p.Name
		if err := os.MkdirAll(newPath, 0755); err != nil {
			return err
		}
		*p.Path = newPath
	}
	// init
	// create project and init
	progressCh <- "Running Project initialization"
	if err := domain.InitProject(uc.Runner, *p.Path, p.ModuleName); err != nil {
		return err
	}

	progressCh <- "Running git initialization"
	// init git or not
	if err := domain.InitGit(uc.Runner, p.Path, p.Git); err != nil {
		return err
	}
	// run each init sequentially (single thread)
	steps := []initStep{
		// config
		{progress: "Running framework generation", run: func() error {
			return domain.InitFramework(p.Path, &p.Framework)
		}},
		{progress: "Running env config generation", run: func() error {
			return domain.InitENVConfig(p.Path)
		}},
		{progress: "Running zerolog config generation", run: func() error {
			return domain.InitZerologConfig(p.Path)
		}},
		{progress: "Running DB config generation", run: func() error {
			return domain.InitDBConfig(p.Path, &p.Database)
		}},
		{progress: "Running mq config generation", run: func() error {
			return domain.InitMQConfig(p)
		}},
		{progress: "Running in-memory store config generation", run: func() error {
			return domain.InitInMemoryConfig(p.Path, &p.InMemory)
		}},
		{progress: "Running object config generation", run: func() error {
			return domain.InitObjectStorageConfig(p.Path, &p.ObjectStorage)
		}},

		// infra
		{progress: "Running querier infra generation", run: func() error {
			return domain.InitQuerierInfra(p.Path, &p.Database)
		}},
		{progress: "Running in-memory infra generation", run: func() error {
			return domain.InitInMemoryInfra(p.Path, &p.InMemory)
		}},
		{progress: "Running mq infra generation", run: func() error {
			return domain.InitMQinfra(p)
		}},
		{progress: "Running object infra generation", run: func() error {
			return domain.InitObjectStorageInfra(p.Path, &p.ObjectStorage)
		}},

		// domain
		{progress: "Running dto example generation", run: func() error {
			return domain.InitDTODomain(p.Path)
		}},
		{progress: "Running repository example generation", run: func() error {
			return domain.InitRepositoryDomain(p.Path, p.ModuleName)
		}},
		{progress: "Running service example generation", run: func() error {
			return domain.InitServiceDomain(p.Path, p.ModuleName)
		}},
		{progress: "Running handler example generation", run: func() error {
			return domain.InitHandlerDomain(p.Path, &p.Framework, p.ModuleName)
		}},
		{progress: "Running http handler example generation", run: func() error {
			return domain.InitHTTPHandlerDomain(p.Path, &p.Framework)
		}},
		{progress: "Running pkg example generation", run: func() error {
			return domain.InitPkgExample(p.Path)
		}},

		// cmd
		{progress: "Running dependency injection file generation", run: func() error {
			return domain.InitDependencyInjection(p)
		}},
		{progress: "Running bootstraper file generation", run: func() error {
			return domain.InitBootStrap(p.Path, p.ModuleName)
		}},
		{progress: "Running server file generation", run: func() error {
			return domain.InitServer(p)
		}},
		{progress: "Running main file generation", run: func() error {
			return domain.InitMain(p.Path, p.ModuleName)
		}},

		// global config
		{progress: "Running init air config", run: func() error {
			return domain.InitAirConfig(uc.Runner, p.Path, p.Air)
		}},
		{progress: "Running config.local.yaml generation", run: func() error {
			return domain.InitYamlConfig(p)
		}},
		{progress: "Running .gitignore file generation", run: func() error {
			return domain.InitGitignoreConfig(p.Path)
		}},
		{progress: "Running Dockerfile file generation", run: func() error {
			return domain.InitDockerFileConfig(p.Path, p.Name)
		}},
		{progress: "Running docker-compose.yml file generation", run: func() error {
			return domain.InitDockerComposeConfig(p)
		}},
		{progress: "Running mq config file generation", run: func() error {
			return domain.InitMQConfigFile(p)
		}},
		{progress: "Running cache config file generation", run: func() error {
			return domain.InitInMemoryConfigFile(p)
		}},
		{progress: "Running object storage config file generation", run: func() error {
			return domain.InitObjectStorageConfigFile(p)
		}},
		{progress: "Running .env.example file generation", run: func() error {
			return domain.InitDotEnvExampleConfig(p.Path)
		}},
		{progress: "Running readme.md file generation", run: func() error {
			return domain.InitReadme(p.Path)
		}},
		{progress: "Running version file generation", run: func() error {
			return domain.InitVersion(p.Path)
		}},
		{progress: "Running build script file generation", run: func() error {
			return domain.InitBuildScript(p.Path, p.ModuleName)
		}},
		{progress: "Running binary build script file generation", run: func() error {
			return domain.InitBinaryBuildScript(p.Path, p.Name)
		}},
		{progress: "Running Makefile file generation", run: func() error {
			return domain.InitMakefile(p.Path)
		}},
	}

	for _, s := range steps {
		progressCh <- s.progress
		if err := s.run(); err != nil {
			return err
		}
	}
	return uc.finalize(*p.Path, progressCh)
}

// finalize runs toolchain commands needing network. Split from Run's
// file generation so scaffold stays testable offline.
func (uc *InitProjectUseCase) finalize(path string, progressCh chan<- string) error {
	progressCh <- "Running go imports to resolve import"
	if err := uc.Runner.Run("go", []string{"run", "golang.org/x/tools/cmd/goimports@latest", "-w", "."}, path); err != nil {
		return err
	}
	progressCh <- "Running go get ./... to download 3rd party modules"
	if err := uc.Runner.Run("go", []string{"get", "./..."}, path); err != nil {
		return err
	}
	progressCh <- "Running go mod tidy"
	if err := uc.Runner.Run("go", []string{"mod", "tidy"}, path); err != nil {
		return err
	}
	return nil
}
