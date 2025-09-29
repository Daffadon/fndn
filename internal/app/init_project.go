package app

import (
	"os"

	"github.com/daffadon/fndn/internal/domain"
	"github.com/daffadon/fndn/internal/infra"
)

type InitProjectUseCase struct {
	Runner infra.CommandRunner
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
	initFuncs := []func() error{
		// config
		func() error {
			progressCh <- "Running framework generation"
			return domain.InitFramework(uc.Runner, p.Path, &p.Framework)
		},
		func() error {
			progressCh <- "Running env config generation"
			return domain.InitENVConfig(uc.Runner, p.Path)
		},
		func() error {
			progressCh <- "Running zerolog config generation"
			return domain.InitZerologConfig(uc.Runner, p.Path)
		},
		func() error {
			progressCh <- "Running redis config generation"
			return domain.InitRedisConfig(uc.Runner, p.Path)
		},
		func() error {
			progressCh <- "Running nats config generation"
			return domain.InitNatsConfig(uc.Runner, p.Path)
		},
		func() error {
			progressCh <- "Running minio config generation"
			return domain.InitMinioConfig(uc.Runner, p.Path)
		},
		func() error {
			progressCh <- "Running postgresql config generation"
			return domain.InitDBConfig(uc.Runner, p.Path, &p.Database)
		},

		// infra
		func() error {
			progressCh <- "Running querier infra generation"
			return domain.InitQuerierInfra(uc.Runner, p.Path, &p.Database)
		},
		func() error {
			progressCh <- "Running redis infra generation"
			return domain.InitRedisInfra(uc.Runner, p.Path)
		},
		func() error {
			progressCh <- "Running jetstream infra generation"
			return domain.InitJetstreamInfra(uc.Runner, p.Path)
		},
		func() error {
			progressCh <- "Running minio infra generation"
			return domain.InitMinioInfra(uc.Runner, p.Path)
		},

		// domain
		func() error {
			progressCh <- "Running dto example generation"
			return domain.InitDTODomain(uc.Runner, p.Path)
		},
		func() error {
			progressCh <- "Running repository example generation"
			return domain.InitRepositoryDomain(uc.Runner, p.Path, p.ModuleName)
		},
		func() error {
			progressCh <- "Running service example generation"
			return domain.InitServiceDomain(uc.Runner, p.Path)
		},
		func() error {
			progressCh <- "Running handler example generation"
			return domain.InitHandlerDomain(uc.Runner, p.Path, &p.Framework)
		},
		func() error {
			progressCh <- "Running http handler example generation"
			return domain.InitHTTPHandlerDomain(uc.Runner, p.Path, &p.Framework)
		},
		func() error {
			progressCh <- "Running pkg example generation"
			return domain.InitPkgExample(uc.Runner, p.Path)
		},

		// cmd
		func() error {
			progressCh <- "Running dependency injection file generation"
			return domain.InitDependencyInjection(uc.Runner, p.Path)
		},
		func() error {
			progressCh <- "Running bootstraper file generation"
			return domain.InitBootStrap(uc.Runner, p.Path)
		},
		func() error {
			progressCh <- "Running server file generation"
			return domain.InitServer(uc.Runner, p)
		},
		func() error {
			progressCh <- "Running main file generation"
			return domain.InitMain(uc.Runner, p.Path)
		},

		// global config
		func() error {
			progressCh <- "Running init air config"
			return domain.InitAirConfig(uc.Runner, p.Path, p.Air)
		},
		func() error {
			progressCh <- "Running config.local.yaml generation"
			return domain.InitYamlConfig(uc.Runner, p)
		},
		func() error {
			progressCh <- "Running .gitignore file generation"
			return domain.InitGitignoreConfig(uc.Runner, p.Path)
		},
		func() error {
			progressCh <- "Running Dockerfile file generation"
			return domain.InitDockerFileConfig(uc.Runner, p.Path, p.Name)
		},
		func() error {
			progressCh <- "Running docker-compose.yml file generation"
			return domain.InitDockerComposeConfig(uc.Runner, p)
		},
		func() error {
			progressCh <- "Running nats-server.conf file generation"
			return domain.InitNatsConfigFile(uc.Runner, p.Path)
		},
		func() error {
			progressCh <- "Running .env.example file generation"
			return domain.InitDotEnvExampleConfig(uc.Runner, p.Path)
		},
		func() error {
			progressCh <- "Running readme.md file generation"
			return domain.InitReadme(uc.Runner, p.Path)
		},
		func() error {
			progressCh <- "Running version file generation"
			return domain.InitVersion(uc.Runner, p.Path)
		},
		func() error {
			progressCh <- "Running build script file generation"
			return domain.InitBuildScript(uc.Runner, p.Path, p.ModuleName)
		},
		func() error {
			progressCh <- "Running binary build script file generation"
			return domain.InitBinaryBuildScript(uc.Runner, p.Path, p.Name)
		},
		func() error {
			progressCh <- "Running Makefile file generation"
			return domain.InitMakefile(uc.Runner, p.Path)
		},
	}

	for _, f := range initFuncs {
		if err := f(); err != nil {
			return err
		}
	}
	progressCh <- "Running go imports to resolve import"
	if err := uc.Runner.Run("go", []string{"run", "golang.org/x/tools/cmd/goimports@latest", "-w", "."}, *p.Path); err != nil {
		return err
	}
	progressCh <- "Running go get -u ./... to download 3rd party modules"
	if err := uc.Runner.Run("go", []string{"get", "-u", "./..."}, *p.Path); err != nil {
		return err
	}
	progressCh <- "Running go mod tidy"
	if err := uc.Runner.Run("go", []string{"mod", "tidy"}, *p.Path); err != nil {
		return err
	}

	return nil
}
