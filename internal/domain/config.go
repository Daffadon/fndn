package domain

import (
	"errors"

	"github.com/daffadon/fndn/internal/pkg"
	config_template "github.com/daffadon/fndn/internal/template/config"
	"github.com/daffadon/fndn/internal/template/readme"
)

var yamlDBConfigs = map[string]string{
	"postgresql": config_template.PostresqlYamlConfigTemplate,
	"mariadb":    config_template.MariaDBYamlConfigTemplate,
	"clickhouse": config_template.ClickHouseYamlConfigTemplate,
	"mongodb":    config_template.MongoDBYamlConfigTemplate,
	"ferretdb":   config_template.FerretDBYamlConfigTemplate,
	"neo4j":      config_template.Neo4JYamlConfigTemplate,
}

var yamlCacheConfigs = map[string]string{
	"redis":     config_template.RedisYamlConfigTemplate,
	"valkey":    config_template.ValkeyYamlConfigTemplate,
	"dragonfly": config_template.DragonFlyYamlConfigTemplate,
	"redict":    config_template.RedictYamlConfigTemplate,
}

var yamlOSConfigs = map[string]string{
	"rustfs":    config_template.RustfsYamlConfigTemplate,
	"seaweedfs": config_template.SeaweedfsYamlConfigTemplate,
	"minio":     config_template.MinioYamlConfigTemplate,
}

func InitENVConfig(path *string) error {
	if path != nil {
		folderName := "/config/env"
		fileName := folderName + "/env.go"
		if err := pkg.GoFileGenerator(path, folderName, fileName, config_template.ENVConfigTemplate); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitYamlConfig(p *Project) error {
	if p.Path != nil {
		folderName := ""
		fileName := folderName + "/config.local.yaml"

		s := config_template.YamlConfigMessageTemplate

		db, err := section(yamlDBConfigs, "database", p.Database)
		if err != nil {
			return err
		}
		s += db
		s += config_template.AppYamlConfigTemplate
		cache, err := section(yamlCacheConfigs, "in-memory store", p.InMemory)
		if err != nil {
			return err
		}
		s += cache

		switch p.MQ {
		case "nats":
			s += config_template.NatsYamlConfigTemplate
			s += config_template.JetstreamConfigTemplate
		case "rabbitmq":
			s += config_template.RabbitYamlConfigTemplate
		case "kafka":
			s += config_template.KafkaYamlConfigTemplate
		case "amazon sqs":
			s += config_template.AmazonSQSConfigTemplate
		case None:
			break
		default:
			_, err := lookup(yamlDBConfigs, "message queue", p.MQ)
			return err
		}

		os, err := section(yamlOSConfigs, "object storage", p.ObjectStorage)
		if err != nil {
			return err
		}
		s += os
		s += config_template.ServerYamlConfigTemplate

		if err := pkg.GenericFileGenerator(p.Path, folderName, fileName, s); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitGitignoreConfig(path *string) error {
	if path != nil {
		folderName := ""
		fileName := folderName + "/.gitignore"
		if err := pkg.GenericFileGenerator(path, folderName, fileName, config_template.GitIgnoreConfigTemplate); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitDotEnvExampleConfig(path *string) error {
	if path != nil {
		folderName := ""
		fileName := folderName + "/.env.example"
		if err := pkg.GenericFileGenerator(path, folderName, fileName, config_template.DotENVExampleTemplate); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitReadme(path *string) error {
	if path != nil {
		folderName := ""
		fileName := folderName + "/README.md"
		s, err := readme.CopyReadmeTemplate()
		if err != nil {
			return err
		}
		if err := pkg.GenericFileGenerator(path, folderName, fileName, s); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}

func InitVersion(path *string) error {
	if path != nil {
		folderName := ""
		fileName := folderName + "/VERSION"
		if err := pkg.GenericFileGenerator(path, folderName, fileName, config_template.VersionConfigTemplate); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
func InitBuildScript(path *string, moduleName string) error {
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
			return err
		}
		if err := pkg.GenericFileGenerator(path, folderName, fileName, c); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
func InitBinaryBuildScript(path *string, projectName string) error {
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
			return err
		}
		if err := pkg.GenericFileGenerator(path, folderName, fileName, c); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
func InitMakefile(path *string) error {
	if path != nil {
		folderName := ""
		fileName := folderName + "/Makefile"
		if err := pkg.GenericFileGenerator(path, folderName, fileName, config_template.MakefileConfigTemplate); err != nil {
			return err
		}
		return nil
	}
	return errors.New("path is nil")
}
