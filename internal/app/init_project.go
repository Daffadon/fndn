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
	// create project and init
	if err := domain.InitProject(uc.Runner, p.Path, p.ModuleName); err != nil {
		return err
	}
	// init git or not
	if err := domain.InitGit(uc.Runner, &p.Path, p.Git); err != nil {
		return err
	}

	if err := domain.InitGin(uc.Runner, &p.Path); err != nil {
		return err
	}
	// init gin
	if err := uc.Runner.Run("go", []string{"mod", "tidy"}, p.Path); err != nil {
		return err
	}
	// which
	return nil
}
