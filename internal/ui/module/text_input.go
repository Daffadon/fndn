package module

import (
	btxt "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/daffadon/fndn/internal/types"
)

type TextInput struct {
	ti btxt.Model
}

func NewTextInput(placeholder, initial string) *TextInput {
	ti := btxt.New()
	ti.Placeholder = placeholder
	ti.SetValue(initial)
	ti.Prompt = ""
	return &TextInput{ti: ti}
}

func (i *TextInput) Update(msg tea.Msg) (types.Input, tea.Cmd) {
	var cmd tea.Cmd
	i.ti, cmd = i.ti.Update(msg)
	return i, cmd
}

func (i *TextInput) View() string { return i.ti.View() }
func (i *TextInput) Value() any   { return i.ti.Value() }
func (i *TextInput) Focus()       { i.ti.Focus() }
func (i *TextInput) Blur()        { i.ti.Blur() }
