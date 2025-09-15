package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/daffadon/fndn/internal/app"
	"github.com/daffadon/fndn/internal/domain"
)

type model struct {
	textInput textinput.Model
	useCase   *app.InitProjectUseCase
	done      bool
	err       error
}

func NewModel(uc *app.InitProjectUseCase) model {
	ti := textinput.New()
	ti.Placeholder = "github.com/you/project"
	ti.Focus()

	return model{
		textInput: ti,
		useCase:   uc,
	}
}

func (m model) Init() tea.Cmd { return textinput.Blink }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			p := domain.Project{ModuleName: m.textInput.Value()}
			m.err = m.useCase.Run(p)
			m.done = true
			return m, tea.Quit
		case "ctrl+c", "esc":
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.done {
		if m.err != nil {
			return fmt.Sprintf("❌ Failed: %v\n", m.err)
		}
		return fmt.Sprintf("✅ go mod init %s done!\n", m.textInput.Value())
	}
	return fmt.Sprintf("Enter your Go module name:\n\n%s\n\n(press Enter to confirm)", m.textInput.View())
}
