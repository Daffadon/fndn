package domain

import (
	"github.com/daffadon/fndn/internal/infra"
)

type Project struct {
	ModuleName string
	Name       string
	Path       *string
	Framework  string
	Database   string
	MQ         string
	Git        bool
	Air        bool
}

func InitProject(i infra.CommandRunner, path, moduleName string) error {
	return i.Run("go", []string{"mod", "init", moduleName}, path)
}

func InitGit(i infra.CommandRunner, path *string, isInit bool) error {
	if isInit {
		return i.Run("git", []string{"init"}, *path)
	}
	return nil
}
