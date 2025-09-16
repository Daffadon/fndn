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
	spinner   spinner.Model
	stopwatch *dto.StopwatchModel
	useCase   *app.InitProjectUseCase
	err       error
	current   int
	targetDir string

	loading bool
	done    bool
}
