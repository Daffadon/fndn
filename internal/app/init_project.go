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
		{"Running framework generation", func() error {
			return domain.InitFramework(p.Path, &p.Framework)
		}},
		{"Running env config generation", func() error {
			return domain.InitENVConfig(p.Path)
		}},
		{"Running zerolog config generation", func() error {
			return domain.InitZerologConfig(p.Path)
		}},
		{"Running DB config generation", func() error {
			return domain.InitDBConfig(p.Path, &p.Database)
		}},
		{"Running mq config generation", func() error {
			return domain.InitMQConfig(p)
		}},
		{"Running in-memory store config generation", func() error {
			return domain.InitInMemoryConfig(p.Path, &p.InMemory)
		}},
		{"Running object config generation", func() error {
			return domain.InitObjectStorageConfig(p.Path, &p.ObjectStorage)
		}},

		// infra
		{"Running querier infra generation", func() error {
			return domain.InitQuerierInfra(p.Path, &p.Database)
		}},
		{"Running in-memory infra generation", func() error {
			return domain.InitInMemoryInfra(p.Path, &p.InMemory)
		}},
		{"Running mq infra generation", func() error {
			return domain.InitMQinfra(p)
		}},
		{"Running object infra generation", func() error {
			return domain.InitObjectStorageInfra(p.Path, &p.ObjectStorage)
		}},

		// domain
		{"Running dto example generation", func() error {
			return domain.InitDTODomain(p.Path)
		}},
		{"Running repository example generation", func() error {
			return domain.InitRepositoryDomain(p.Path, p.ModuleName)
		}},
		{"Running service example generation", func() error {
			return domain.InitServiceDomain(p.Path)
		}},
		{"Running handler example generation", func() error {
			return domain.InitHandlerDomain(p.Path, &p.Framework)
		}},
		{"Running http handler example generation", func() error {
			return domain.InitHTTPHandlerDomain(p.Path, &p.Framework)
		}},
		{"Running pkg example generation", func() error {
			return domain.InitPkgExample(p.Path)
		}},

		// cmd
		{"Running dependency injection file generation", func() error {
			return domain.InitDependencyInjection(p)
		}},
		{"Running bootstraper file generation", func() error {
			return domain.InitBootStrap(p.Path)
		}},
		{"Running server file generation", func() error {
			return domain.InitServer(p)
		}},
		{"Running main file generation", func() error {
			return domain.InitMain(p.Path)
		}},

		// global config
		{"Running init air config", func() error {
			return domain.InitAirConfig(uc.Runner, p.Path, p.Air)
		}},
		{"Running config.local.yaml generation", func() error {
			return domain.InitYamlConfig(p)
		}},
		{"Running .gitignore file generation", func() error {
			return domain.InitGitignoreConfig(p.Path)
		}},
		{"Running Dockerfile file generation", func() error {
			return domain.InitDockerFileConfig(p.Path, p.Name)
		}},
		{"Running docker-compose.yml file generation", func() error {
			return domain.InitDockerComposeConfig(p)
		}},
		{"Running mq config file generation", func() error {
			return domain.InitMQConfigFile(p)
		}},
		{"Running cache config file generation", func() error {
			return domain.InitInMemoryConfigFile(p)
		}},
		{"Running object storage config file generation", func() error {
			return domain.InitObjectStorageConfigFile(p)
		}},
		{"Running .env.example file generation", func() error {
			return domain.InitDotEnvExampleConfig(p.Path)
		}},
		{"Running readme.md file generation", func() error {
			return domain.InitReadme(p.Path)
		}},
		{"Running version file generation", func() error {
			return domain.InitVersion(p.Path)
		}},
		{"Running build script file generation", func() error {
			return domain.InitBuildScript(p.Path, p.ModuleName)
		}},
		{"Running binary build script file generation", func() error {
			return domain.InitBinaryBuildScript(p.Path, p.Name)
		}},
		{"Running Makefile file generation", func() error {
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
	progressCh <- "Running go get -u ./... to download 3rd party modules"
	if err := uc.Runner.Run("go", []string{"get", "-u", "./..."}, path); err != nil {
		return err
	}
	progressCh <- "Running go mod tidy"
	if err := uc.Runner.Run("go", []string{"mod", "tidy"}, path); err != nil {
		return err
	}
	return nil
}
