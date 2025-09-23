package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/daffadon/fndn/internal/domain"
	"github.com/daffadon/fndn/internal/pkg"
	"github.com/daffadon/fndn/internal/ui/helper"
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
	width := m.width
	if width <= 0 {
		width = 80 // fallback until WindowSizeMsg arrives
	}
	line := fmt.Sprintf(
		"%s %s | elapsed: %.1fs",
		m.spinner.View(),
		m.logs,
		m.stopwatch.Elapsed().Seconds(),
	)
	return helper.PadOrTruncate(line, width)
}

func (m *model) viewDone() string {
	if m.err != nil {
		return fmt.Sprintf("❌ Failed: %v\n", m.err)
	}
	pn := style.BlueStyle.Render(m.steps[0].Input.Value().(string))
	return fmt.Sprintf(
		"project %s has been generated!\nelapsed time: %.1fs\nCheck Readme for further step after initialization\n",
		pn,
		m.stopwatch.Elapsed().Seconds(),
	)
}

func (m *model) viewStep() string {
	var s string
	s += "\nThis will create a new Go module and scaffold a basic clean-code architecture\n"

	total := len(m.steps)
	label := style.BlueStyle.Render(fmt.Sprintf("Step %d/%d", m.current+1, total))

	s += fmt.Sprintf("%s\n", label)
	switch m.current {
	case 0:
		s += style.BlueStyle.Render("enter your module name\n")
	case 1:
		s += style.BlueStyle.Render("would you init git also?\n")
	case 2:
		s += style.BlueStyle.Render("would you use air for hot reload?\n")
	}

	s += "\n"
	content := m.steps[m.current].Input.View()
	arrow := style.ArrowStyle.Render("> ")

	s += fmt.Sprintf("%s%s\n\n", arrow, content)

	if m.err != nil {
		s += style.ErrorStyle.Render(fmt.Sprintf("\n⚠️  %v\n", m.err))
	}

	if m.current < total-1 {
		s += "\n(Enter to continue; Left/Shift+Tab to go back; Esc to cancel)\n"
	} else {
		s += "\n(Enter to scaffold; Left/Shift+Tab to go back; Esc to cancel)\n"
	}

	return s
}

func (m *model) submit() tea.Cmd {
	m.stopwatch.Start()
	moduleName := m.steps[0].Input.Value().(string)

	initGit := false
	if v, ok := m.steps[1].Input.Value().(bool); ok {
		initGit = v
	}
	initAir := false
	if v, ok := m.steps[2].Input.Value().(bool); ok {
		initAir = v
	}

	name := pkg.LastSegment(moduleName)
	project := domain.Project{
		ModuleName: moduleName,
		Name:       name,
		Path:       m.targetDir,
		Git:        initGit,
		Air:        initAir,
	}

	m.loading = true
	m.err = nil
	return tea.Batch(m.waitForProgress(), m.runInitProject(project), m.spinner.Tick)
}

func (m *model) runInitProject(p domain.Project) tea.Cmd {
	return func() tea.Msg {
		go func() {
			err := m.useCase.Run(&p, m.progressCh)
			close(m.progressCh)
			m.errCh <- err
			close(m.errCh)
		}()
		return m.waitForProgress()
	}
}

func (m *model) waitForProgress() tea.Cmd {
	return func() tea.Msg {
		msg, ok := <-m.progressCh
		if !ok {
			err := <-m.errCh
			return initFinishedMsg{err: err}
		}
		return progressMsg(msg)
	}
}
