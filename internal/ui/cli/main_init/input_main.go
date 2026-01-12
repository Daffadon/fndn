package main_init

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/daffadon/fndn/assets"
	"github.com/daffadon/fndn/internal/app"
	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/ui/dto"
	"github.com/daffadon/fndn/internal/ui/module"
	"github.com/daffadon/fndn/internal/ui/style"
)

func newModel(uc *app.InitProjectUseCase, targetDir string) model {
	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))

	steps := []dto.Step{
		{
			Label: "Go module name",
			Input: module.NewTextInput("Module name (e.g. github.com/you/project)", "github.com/your-name/project"),
			Validate: func(v any) error {
				s, ok := v.(string)
				if !ok || strings.TrimSpace(s) == "" {
					return fmt.Errorf("module name cannot be empty")
				}
				return nil
			},
		},
		{
			Label:    "Initialize Git repo",
			Input:    module.NewCheckbox("Init git?", true),
			Validate: nil,
		},
		{
			Label:    "Initialize hot reload with air",
			Input:    module.NewCheckbox("Init air?", true),
			Validate: nil,
		},
		{
			Label:    "Generated set of module",
			Input:    module.NewRadioButton([]string{"Default", "Custom"}, 0),
			Validate: nil,
		},
		{
			Label:    "Framework",
			Input:    module.NewRadioButton([]string{"Gin", "Chi", "Echo", "Fiber", "Gorilla/mux"}, 0),
			Validate: nil,
		},
		{
			Label:    "Database",
			Input:    module.NewRadioButton([]string{"Postgresql", "MariaDB", "ClickHouse", "MongoDB", "FerretDB", "Neo4j"}, 0),
			Validate: nil,
		},
		{
			Label:    "Message Queue",
			Input:    module.NewRadioButton([]string{"Nats", "RabbitMQ", "Kafka", "Amazon SQS"}, 0),
			Validate: nil,
		},
		{
			Label:    "In-memory Store",
			Input:    module.NewRadioButton([]string{"Redis", "Valkey", "Dragonfly", "Redict"}, 0),
			Validate: nil,
		},
		{
			Label:    "Object Storage",
			Input:    module.NewRadioButton([]string{"RustFS", "SeaweedFS", "MinIO"}, 0),
			Validate: nil,
		},
	}

	steps[0].Input.Focus()

	return model{
		steps:      steps,
		current:    0,
		spinner:    sp,
		useCase:    uc,
		targetDir:  targetDir,
		stopwatch:  &dto.StopwatchModel{},
		progressCh: make(chan string),
		errCh:      make(chan error, 1),
	}
}

func RunModuleInput(targetDir string) error {
	uc := &app.InitProjectUseCase{Runner: infra.NewCommandRunner()}
	p := tea.NewProgram(newModel(uc, targetDir))
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "shift+tab":
			m.prevStep()
		case "enter":
			if m.loading {
				return m, nil
			}
			step := m.steps[m.current]
			if step.Validate != nil {
				if err := step.Validate(step.Input.Value()); err != nil {
					m.err = err
					return m, nil
				}
			}
			if m.current < len(m.steps)-1 {
				if m.current == 3 && m.steps[3].Input.Value().(string) == "Default" {
					return m, m.submit()
				}
				m.nextStep()
			} else {
				return m, m.submit()
			}
		}

	case initFinishedMsg:
		m.loading = false
		m.done = true
		m.stopwatch.Stop()
		m.err = msg.err
		return m, tea.Quit

	case progressMsg:
		now := time.Now()
		if now.Sub(m.lastProgress) > 50*time.Millisecond {
			m.logs = string(msg)
			m.lastProgress = now
		}
		return m, m.waitForProgress()
	}

	if m.loading {
		var spCmd tea.Cmd
		m.spinner, spCmd = m.spinner.Update(msg)
		return m, spCmd
	}

	in, cmd := m.steps[m.current].Input.Update(msg)
	m.steps[m.current].Input = in
	return m, cmd
}

func (m model) View() string {
	logo := style.BlueStyle.Render(assets.Logo)
	switch {
	case m.loading:
		return logo + "\n" + m.viewLoading()
	case m.done:
		return logo + "\n" + m.viewDone()
	default:
		return logo + "\n" + m.viewStep()
	}
}
