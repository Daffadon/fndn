package app

import (
	"github.com/daffadon/fndn/internal/domain"
	"github.com/daffadon/fndn/internal/infra"
)

type InitProjectUseCase struct {
	Runner infra.CommandRunner
}

func (uc *InitProjectUseCase) Run(p domain.Project) error {
	return uc.Runner.Run("go", "mod", "init", p.ModuleName)
}
