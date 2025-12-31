package ui

import (
	"fmt"
	"strings"

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

	switch m.current {
	case 0:
		s += style.BlueStyle.Render("enter your module name\n")
	case 1:
		s += style.BlueStyle.Render("would you init git also?\n")
	case 2:
		s += style.BlueStyle.Render("would you use air for hot reload?\n")
	case 3:
		s += style.BlueStyle.Render("which set would you generate?\n")
	case 4:
		s += style.BlueStyle.Render("which framework would you use?\n")
	case 5:
		s += style.BlueStyle.Render("which database would you working on?\n")
	case 6:
		s += style.BlueStyle.Render("which message queue would you connect?\n")
	case 7:
		s += style.BlueStyle.Render("which in-memory store would you use?\n")
	case 8:
		s += style.BlueStyle.Render("which object storage whould you choose?\n")
	}

	s += "\n"
	content := m.steps[m.current].Input.View()

	s += fmt.Sprintf("%s\n\n", content)

	if m.err != nil {
		s += style.ErrorStyle.Render(fmt.Sprintf("\n⚠️  %v\n", m.err))
	}

	if m.current == 3 && m.steps[3].Input.Value().(string) == "Default" {
		s += "(Enter to scaffold; Left/Shift+Tab to go back; Esc to cancel)\n"
	} else if m.current < total-1 {
		s += "(Enter to continue; Left/Shift+Tab to go back; Esc to cancel)\n"
	} else if m.current == total-1 {
		s += "(Enter to scaffold; Left/Shift+Tab to go back; Esc to cancel)\n"
	}

	return s
}

func (m *model) submit() tea.Cmd {
	m.stopwatch.Start()
	project := domain.Project{
		ModuleName:    "",
		Name:          "",
		Path:          nil,
		Git:           false,
		Air:           false,
		Framework:     "gin",
		Database:      "postgresql",
		MQ:            "nats",
		InMemory:      "redis",
		ObjectStorage: "rustfs",
	}
	project.Path = &m.targetDir

	moduleName := m.steps[0].Input.Value().(string)
	project.ModuleName = moduleName

	name := pkg.LastSegment(moduleName)
	project.Name = name

	if v, ok := m.steps[1].Input.Value().(bool); ok {
		project.Git = v
	}
	if v, ok := m.steps[2].Input.Value().(bool); ok {
		project.Air = v
	}
	if v, ok := m.steps[3].Input.Value().(string); ok {
		if strings.ToLower(v) != "default" {
			if v, ok := m.steps[4].Input.Value().(string); ok {
				project.Framework = strings.ToLower(v)
			}
			if v, ok := m.steps[5].Input.Value().(string); ok {
				project.Database = strings.ToLower(v)
			}
			if v, ok := m.steps[6].Input.Value().(string); ok {
				project.MQ = strings.ToLower(v)
			}
			if v, ok := m.steps[7].Input.Value().(string); ok {
				project.InMemory = strings.ToLower(v)
			}
			if v, ok := m.steps[8].Input.Value().(string); ok {
				project.ObjectStorage = strings.ToLower(v)
			}
		}
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
