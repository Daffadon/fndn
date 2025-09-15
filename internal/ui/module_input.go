package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/daffadon/fndn/internal/app"
	"github.com/daffadon/fndn/internal/domain"
	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
)

type model struct {
	textInput textinput.Model
	spinner   spinner.Model
	useCase   *app.InitProjectUseCase
	targetDir string
	loading   bool
	done      bool
	err       error
	startTime time.Time
	elapsed   time.Duration
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.spinner.Tick)
}

func newModel(uc *app.InitProjectUseCase, targetDir string) model {
	ti := textinput.New()
	ti.SetValue("github.com/your-name/project")
	ti.Focus()

	sp := spinner.New()
	sp.Spinner = spinner.Dot

	return model{
		textInput: ti,
		spinner:   sp,
		useCase:   uc,
		elapsed:   0,
		targetDir: targetDir,
	}
}

func RunModuleInput(targetDir string) error {
	p := tea.NewProgram(newModel(&app.InitProjectUseCase{Runner: infra.NewCommandRunner()}, targetDir))
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.loading = true
			m.startTime = time.Now()

			name := pkg.LastSegment(m.textInput.Value())

			p := domain.Project{
				ModuleName: m.textInput.Value(),
				Name:       name,
				Path:       m.targetDir,
			}

			m.err = m.useCase.Run(&p)

			m.done = true
			m.loading = false
			m.elapsed = time.Since(m.startTime)

			return m, tea.Quit
		case "ctrl+c", "esc":
			return m, tea.Quit
		}
	}
	if m.loading {
		var spCmd tea.Cmd
		m.spinner, spCmd = m.spinner.Update(msg)
		m.elapsed = time.Since(m.startTime)
		cmds = append(cmds, spCmd)
	}
	m.textInput, _ = m.textInput.Update(msg)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.loading {
		return fmt.Sprintf("Initializing module...\nElapsed: %.1fs\n%s", m.elapsed.Seconds(), m.spinner.View())
	}
	if m.done {
		if m.err != nil {
			return fmt.Sprintf("❌ Failed: %v\nElapsed: %.1fs\n", m.err, m.elapsed.Seconds())
		}
		return fmt.Sprintf("✅ go mod init %s done!\nElapsed: %.1fs\n", m.textInput.Value(), m.elapsed.Seconds())
	}
	return fmt.Sprintf("Enter your Go module name:\n\n%s\n\n(press Enter to confirm)", m.textInput.View())
}
