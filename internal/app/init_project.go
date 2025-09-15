package app

import (
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
	return uc.Runner.Run("go", []string{"mod", "init", p.ModuleName}, p.Path)
}
