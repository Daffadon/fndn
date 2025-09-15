package types

import tea "github.com/charmbracelet/bubbletea"

type Input interface {
	Update(msg tea.Msg) (Input, tea.Cmd)
	View() string
	Value() any
	Focus()
	Blur()
}
