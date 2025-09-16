package ui

import (
	"fmt"
	"strings"

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

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.spinner.Tick)
}

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
	}

	steps[0].Input.Focus()

	return model{
		steps:     steps,
		current:   0,
		spinner:   sp,
		useCase:   uc,
		targetDir: targetDir,
		stopwatch: &dto.StopwatchModel{},
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

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "tab", "right":
			m.nextStep()
		case "shift+tab", "left":
			m.prevStep()
		case "enter":
			step := m.steps[m.current]
			if step.Validate != nil {
				if err := step.Validate(step.Input.Value()); err != nil {
					m.err = err
					return m, nil
				}
			}
			if m.current < len(m.steps)-1 {
				m.nextStep()
			} else {
				return m, m.submit()
			}
		}

	case initFinishedMsg:
		m.loading = false
		m.done = true
		m.err = msg.err
		return m, tea.Quit
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
