package ui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/daffadon/fndn/internal/domain"
	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/pkg"
	"github.com/daffadon/fndn/internal/ui/style"
)

func (m *model) nextStep() {
	if m.current < len(m.steps)-1 {
		m.steps[m.current].Input.Blur()
		m.current++
		m.steps[m.current].Input.Focus()
	}
}

func (m *model) prevStep() {
	if m.current > 0 {
		m.steps[m.current].Input.Blur()
		m.current--
		m.steps[m.current].Input.Focus()
	}
}

func (m *model) viewLoading() string {
	return fmt.Sprintf("Initializing project...\nElapsed: %.1fs\n\n%s\n",
		m.elapsed.Seconds(), m.spinner.View())
}

func (m *model) viewDone() string {
	if m.err != nil {
		return fmt.Sprintf("❌ Failed: %v\nElapsed: %.1fs\n", m.err, m.elapsed.Seconds())
	}
	// Use steps[0] because module name is first input
	return fmt.Sprintf("✅ go mod init %s done!\nElapsed: %.1fs\n",
		m.steps[0].Input.Value().(string), m.elapsed.Seconds())
}

func (m *model) viewStep() string {
	var s string
	s += "\nThis will create a new Go module and scaffold a basic clean-code architecture\n"

	total := len(m.steps)
	label := style.StepLabelStyle.Render(fmt.Sprintf("Step %d/%d", m.current+1, total))
	content := m.steps[m.current].Input.View()
	arrow := style.ArrowStyle.Render("> ")

	s += fmt.Sprintf("%s\n\n%s%s\n\n", label, arrow, content)

	if m.err != nil {
		s += style.ErrorStyle.Render(fmt.Sprintf("\n⚠️  %v\n", m.err))
	}

	// hints
	if m.current < total-1 {
		s += "\n(Enter to continue; Left/Shift+Tab to go back; Esc to cancel)\n"
	} else {
		s += "\n(Enter to scaffold; Left/Shift+Tab to go back; Esc to cancel)\n"
	}

	return s
}

func (m *model) submit() tea.Cmd {
	moduleName := m.steps[0].Input.Value().(string)

	initGit := false
	if v, ok := m.steps[1].Input.Value().(bool); ok {
		initGit = v
	}

	// derive project name and path
	name := pkg.LastSegment(moduleName)
	project := domain.Project{
		ModuleName: moduleName,
		Name:       name,
		Path:       m.targetDir,
	}

	m.loading = true
	m.startTime = time.Now()
	m.err = nil
	return m.runInitProject(project, initGit)
}

func (m *model) runInitProject(p domain.Project, initGit bool) tea.Cmd {
	return func() tea.Msg {
		err := m.useCase.Run(&p)

		if err == nil && initGit {
			runner := infra.NewCommandRunner()
			_ = runner.Run("git", []string{"init"}, p.Path)
		}
		return initFinishedMsg{err: err}
	}
}
