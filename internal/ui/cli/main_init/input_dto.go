package main_init

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/daffadon/fndn/internal/app"
	"github.com/daffadon/fndn/internal/ui/dto"
)

type initFinishedMsg struct {
	err error
}
type model struct {
	steps        []dto.Step
	spinner      spinner.Model
	stopwatch    *dto.StopwatchModel
	useCase      *app.InitProjectUseCase
	width        int
	current      int
	lastProgress time.Time
	progressCh   chan string
	errCh        chan error
	logs         string
	err          error
	targetDir    string

	loading bool
	done    bool
}

type progressMsg string
