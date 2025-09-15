package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/daffadon/fndn/internal/app"
	"github.com/daffadon/fndn/internal/types"
)

type Step struct {
	Label    string
	Input    types.Input
	Validate func(value any) error
}
type initFinishedMsg struct {
	err error
}
type model struct {
	steps     []Step
	current   int
	spinner   spinner.Model
	useCase   *app.InitProjectUseCase
	targetDir string

	loading   bool
	done      bool
	err       error
	startTime time.Time
	elapsed   time.Duration
}
