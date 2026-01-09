package main_generate

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/daffadon/fndn/assets"
	"github.com/daffadon/fndn/internal/app"
	"github.com/daffadon/fndn/internal/domain"
	"github.com/daffadon/fndn/internal/ui/dto"
	"github.com/daffadon/fndn/internal/ui/helper"
	"github.com/daffadon/fndn/internal/ui/style"
)

type emitFinishedMsg struct {
	err error
}
type progressMsg string
type GenerateModel struct {
	Steps        []dto.Step
	Spinner      spinner.Model
	Stopwatch    *dto.StopwatchModel
	Ig           *app.InitGenerate
	Width        int
	Current      int
	LastProgress time.Time
	ProgressCh   chan string
	ErrCh        chan error
	Logs         string
	ConfigType   string
	ValueType    string
	Err          error

	Loading bool
	Done    bool
}

func (m GenerateModel) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.Spinner.Tick)
}

func (m GenerateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "enter":
			if m.Loading {
				return m, nil
			}
			step := m.Steps[m.Current]
			if step.Validate != nil {
				if err := step.Validate(step.Input.Value()); err != nil {
					m.Err = err
					return m, nil
				}
			}
			return m, m.submit()
		}

	case emitFinishedMsg:
		m.Loading = false
		m.Done = true
		m.Stopwatch.Stop()
		m.Err = msg.err
		return m, tea.Quit

	case progressMsg:
		now := time.Now()
		if now.Sub(m.LastProgress) > 50*time.Millisecond {
			m.Logs = string(msg)
			m.LastProgress = now
		}
		return m, m.waitForProgress()
	}

	if m.Loading {
		var spCmd tea.Cmd
		m.Spinner, spCmd = m.Spinner.Update(msg)
		return m, spCmd
	}

	in, cmd := m.Steps[m.Current].Input.Update(msg)
	m.Steps[m.Current].Input = in
	return m, cmd
}

func (m GenerateModel) View() string {
	logo := style.BlueStyle.Render(assets.Logo)
	switch {
	case m.Loading:
		return logo + "\n" + m.viewLoading()
	case m.Done:
		return logo + "\n" + m.viewDone()
	default:
		return logo + "\n" + m.viewStep()
	}
}

func (m *GenerateModel) viewLoading() string {
	width := m.Width
	if width <= 0 {
		width = 80 // fallback until WindowSizeMsg arrives
	}
	line := fmt.Sprintf(
		"%s %s | elapsed: %.1fs",
		m.Spinner.View(),
		m.Logs,
		m.Stopwatch.Elapsed().Seconds(),
	)
	return helper.PadOrTruncate(line, width)
}

func (m *GenerateModel) viewDone() string {
	if m.Err != nil {
		return fmt.Sprintf("❌ Failed: %v\n", m.Err)
	}
	di := style.BlueStyle.Render("Don't forget to add the config generated to container.go for dependency injection")
	s := fmt.Sprintf(
		"%s - %s has been generated!\nelapsed time: %.1fs\n%s\n",
		m.ConfigType,
		m.ValueType,
		m.Stopwatch.Elapsed().Seconds(),
		di,
	)
	if m.ConfigType != "framework" {
		s += style.BlueStyle.Render("please take a look at https://github.com/daffadon/fndn/internal/template/common/all_config.yaml.md and search the keyword for additional config.local.yaml to make sure generated config is working\n")
		s += style.BlueStyle.Render("please take a look at https://github.com/daffadon/fndn/internal/template/common/docker-compose.all.md and search the keyword to deploy associated product with generated config\n")
		s += style.BlueStyle.Render("please take a look at https://github.com/daffadon/fndn/internal/template/common/platform_config_file.md and search the keyword to make a file associated to a product deployed (if needed)\n")
	}
	return s
}

func (m *GenerateModel) viewStep() string {
	var s string
	switch m.ConfigType {
	case "framework":
		s += "\nThis will generate a new config for chosen http framework\n"
	}

	s += "\n"
	content := m.Steps[m.Current].Input.View()

	s += fmt.Sprintf("%s\n\n", content)

	if m.Err != nil {
		s += style.ErrorStyle.Render(fmt.Sprintf("\n⚠️  %v\n", m.Err))
	}
	return s
}

func (m *GenerateModel) submit() tea.Cmd {
	m.Stopwatch.Start()
	g := &domain.Generator{
		Type: m.ConfigType,
	}
	if v, ok := m.Steps[0].Input.Value().(string); ok {
		m.ValueType = strings.ToLower(v)
		g.Value = strings.ToLower(v)
	}

	m.Loading = true
	m.Err = nil
	return tea.Batch(m.waitForProgress(), m.runInitProject(g), m.Spinner.Tick)
}

func (m *GenerateModel) runInitProject(g *domain.Generator) tea.Cmd {
	return func() tea.Msg {
		go func() {
			err := m.Ig.Run(g, m.ProgressCh)
			close(m.ProgressCh)
			m.ErrCh <- err
			close(m.ErrCh)
		}()
		return m.waitForProgress()
	}
}

func (m *GenerateModel) waitForProgress() tea.Cmd {
	return func() tea.Msg {
		msg, ok := <-m.ProgressCh
		if !ok {
			err := <-m.ErrCh
			return emitFinishedMsg{err: err}
		}
		return progressMsg(msg)
	}
}
