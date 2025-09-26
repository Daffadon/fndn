package module

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/daffadon/fndn/internal/types"
	"github.com/daffadon/fndn/internal/ui/style"
)

type CheckboxInput struct {
	label   string
	checked bool
	focused bool
}

func NewCheckbox(label string, initial bool) *CheckboxInput {
	return &CheckboxInput{label: label, checked: initial, focused: false}
}

func (c *CheckboxInput) Update(msg tea.Msg) (types.Input, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok && key.String() == " " {
		c.checked = !c.checked
	}
	return c, nil
}

func (c *CheckboxInput) View() string {
	mark := " "
	if c.checked {
		mark = style.BlueStyle.Render("x")
	}

	arrow := ""
	if c.focused {
		arrow = style.ArrowStyle.Render("> ")
	}

	return arrow + fmt.Sprintf("[%s] %s", mark, c.label)
}

func (c *CheckboxInput) Value() any { return c.checked }
func (c *CheckboxInput) Focus()     { c.focused = true }
func (c *CheckboxInput) Blur()      { c.focused = false }
