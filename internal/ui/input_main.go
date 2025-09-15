package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/daffadon/fndn/internal/app"
	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/ui/module"
)

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.spinner.Tick)
}

func newModel(uc *app.InitProjectUseCase, targetDir string) model {
	sp := spinner.New()
	sp.Spinner = spinner.Dot

	steps := []Step{
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
			Validate: nil, // no validation needed
		},
	}

	// Focus the first input
	steps[0].Input.Focus()

	return model{
		steps:     steps,
		current:   0,
		spinner:   sp,
		useCase:   uc,
		targetDir: targetDir,
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
		m.elapsed = time.Since(m.startTime)
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
	switch {
	case m.loading:
		return m.viewLoading()
	case m.done:
		return m.viewDone()
	default:
		return m.viewStep()
	}
}
