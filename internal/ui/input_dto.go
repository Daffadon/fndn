package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/daffadon/fndn/internal/app"
	"github.com/daffadon/fndn/internal/ui/dto"
)

type initFinishedMsg struct {
	err error
}
type model struct {
	steps     []dto.Step
	current   int
	spinner   spinner.Model
	stopwatch *dto.StopwatchModel
	useCase   *app.InitProjectUseCase
	targetDir string

	loading bool
	done    bool
	err     error
}
