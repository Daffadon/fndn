package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	config_template "github.com/daffadon/fndn/internal/template/config"
	"github.com/daffadon/fndn/internal/template/readme"
)

func InitENVConfig(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := "/config/env"
		fileName := folderName + "/env.go"
		if err := pkg.GoFileGenerator(i, path, folderName, fileName, config_template.ENVConfigTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitYamlConfig(i infra.CommandRunner, p *Project) error {
	if p.Path != nil {
		folderName := ""
		fileName := folderName + "/config.local.yaml"

		s := config_template.YamlConfigMessageTemplate

		switch p.Database {
		case "postgresql":
			s += config_template.PostresqlYamlConfigTemplate
		case "mariadb":
			s += config_template.MariaDBYamlConfigTemplate
		case "mongodb":
			s += config_template.MongoDBYamlConfigTemplate
		case "ferretdb":
			s += config_template.FerretDBYamlConfigTemplate
		}
		s += config_template.AppYamlConfigTemplate
		s += config_template.RedisYamlConfigTemplate
		s += config_template.NatsYamlConfigTemplate
		s += config_template.JetstreamConfigTemplate
		s += config_template.MinioYamlConfigTemplate
		s += config_template.ServerYamlConfigTemplate

		if err := pkg.GenericFileGenerator(i, p.Path, folderName, fileName, s); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitGitignoreConfig(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := ""
		fileName := folderName + "/.gitignore"
		if err := pkg.GenericFileGenerator(i, path, folderName, fileName, config_template.GitIgnoreConfigTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitDotEnvExampleConfig(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := ""
		fileName := folderName + "/.env.example"
		if err := pkg.GenericFileGenerator(i, path, folderName, fileName, config_template.DotENVExampleTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitReadme(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := ""
		fileName := folderName + "/README.md"
		s, err := readme.CopyReadmeTemplate()
		if err != nil {
			return err
		}
		if err := pkg.GenericFileGenerator(i, path, folderName, fileName, s); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitVersion(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := ""
		fileName := folderName + "/VERSION"
		if err := pkg.GenericFileGenerator(i, path, folderName, fileName, config_template.VersionConfigTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
func InitBuildScript(i infra.CommandRunner, path *string, moduleName string) error {
	if path != nil {
		folderName := "/script"
		fileName := folderName + "/docker-build.sh"
		st := struct {
			ModuleName string
		}{
			ModuleName: moduleName,
		}
		c, err := pkg.ParseTemplate(config_template.BuildConfigTemplate, st)
		if err != nil {
			log.Fatal(err)
			return err
		}
		if err := pkg.GenericFileGenerator(i, path, folderName, fileName, c); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
func InitBinaryBuildScript(i infra.CommandRunner, path *string, projectName string) error {
	if path != nil {
		folderName := "/script"
		fileName := folderName + "/build-binary.sh"
		st := struct {
			ProjectName string
		}{
			ProjectName: projectName,
		}
		c, err := pkg.ParseTemplate(config_template.BinaryBuildConfigTemplate, st)
		if err != nil {
			log.Fatal(err)
			return err
		}
		if err := pkg.GenericFileGenerator(i, path, folderName, fileName, c); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
func InitMakefile(i infra.CommandRunner, path *string) error {
	if path != nil {
		folderName := ""
		fileName := folderName + "/Makefile"
		if err := pkg.GenericFileGenerator(i, path, folderName, fileName, config_template.MakefileConfigTemplate); err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
