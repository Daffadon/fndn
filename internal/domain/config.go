package domain

import (
	"errors"
	"log"

	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	config_template "github.com/daffadon/fndn/internal/template/config"
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

func InitYamlConfig(i infra.CommandRunner, p *Project, path *string) error {
	if path != nil {
		folderName := ""
		fileName := folderName + "/config.local.yaml"

		s := config_template.YamlConfigMessageTemplate
		s += config_template.PostresqlYamlConfigTemplate
		s += config_template.AppYamlConfigTemplate
		s += config_template.RedisYamlConfigTemplate
		s += config_template.NatsYamlConfigTemplate
		s += config_template.JetstreamConfigTemplate
		s += config_template.MinioYamlConfigTemplate
		s += config_template.ServerYamlConfigTemplate

		if err := pkg.GenericFileGenerator(i, path, folderName, fileName, s); err != nil {
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
