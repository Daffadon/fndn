package cli_storage

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/daffadon/fndn/internal/app"
	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/ui/cli/main_generate"
	"github.com/daffadon/fndn/internal/ui/dto"
	"github.com/daffadon/fndn/internal/ui/module"
)

func newStorageModel(ig *app.InitGenerate) main_generate.GenerateModel {
	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))

	steps := []dto.Step{
		{
			Label:    "Object Storage",
			Input:    module.NewRadioButton([]string{"RustFS", "SeaweedFS", "MinIO"}, 0),
			Validate: nil,
		},
	}

	steps[0].Input.Focus()
	return main_generate.GenerateModel{
		Steps:      steps,
		Current:    0,
		Spinner:    sp,
		Ig:         ig,
		ConfigType: "storage",
		Stopwatch:  &dto.StopwatchModel{},
		ProgressCh: make(chan string),
		ErrCh:      make(chan error, 1),
	}
}

func RunGenerateStorageConfig() error {
	ig := &app.InitGenerate{Runner: infra.NewCommandRunner()}
	p := tea.NewProgram(newStorageModel(ig))
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}
