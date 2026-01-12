package cli_mq

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

func newMQModel(ig *app.InitGenerate) main_generate.GenerateModel {
	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))

	steps := []dto.Step{
		{
			Label:    "Message Queue",
			Input:    module.NewRadioButton([]string{"Nats", "RabbitMQ", "Kafka", "Amazon SQS"}, 0),
			Validate: nil,
		},
	}

	steps[0].Input.Focus()
	return main_generate.GenerateModel{
		Steps:      steps,
		Current:    0,
		Spinner:    sp,
		Ig:         ig,
		ConfigType: "mq",
		Stopwatch:  &dto.StopwatchModel{},
		ProgressCh: make(chan string),
		ErrCh:      make(chan error, 1),
	}
}

func RunGenerateMQConfig() error {
	ig := &app.InitGenerate{Runner: infra.NewCommandRunner()}
	p := tea.NewProgram(newMQModel(ig))
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}
