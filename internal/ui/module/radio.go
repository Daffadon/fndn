package module

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/daffadon/fndn/internal/types"
	"github.com/daffadon/fndn/internal/ui/style"
)

type RadioButton struct {
	options  []string
	selected int
	cursor   int
	focused  bool
}

func NewRadioButton(options []string, initialSelection int) *RadioButton {
	if initialSelection < 0 || initialSelection >= len(options) {
		initialSelection = 0
	}
	return &RadioButton{
		options:  options,
		selected: initialSelection,
		cursor:   initialSelection,
		focused:  false,
	}
}

func (r *RadioButton) Update(msg tea.Msg) (types.Input, tea.Cmd) {
	if !r.focused {
		return r, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if r.cursor > 0 {
				r.cursor--
			}
		case "down", "j":
			if r.cursor < len(r.options)-1 {
				r.cursor++
			}
		case "enter", " ":
			r.selected = r.cursor
		}
	}
	return r, nil
}

func (r *RadioButton) View() string {
	s := ""
	for i, option := range r.options {
		// Show arrow for current cursor position when focused
		cursor := "  "
		if r.cursor == i && r.focused {
			cursor = style.ArrowStyle.Render("> ")
		} else {
			cursor = "  " // Two spaces to match "> " width
		}

		// Show selection state
		checked := " "
		if r.selected == i {
			checked = "●"
		} else {
			checked = "○"
		}

		s += cursor + checked + " " + option + "\n"
	}

	// Add instruction for navigation and selection
	if r.focused {
		s += "\n↑↓ to navigate, Space to select, Enter to choose "
	}

	return s[:len(s)-1] // Remove trailing newline
}

func (r *RadioButton) Value() any {
	if r.selected >= 0 && r.selected < len(r.options) {
		return r.options[r.selected]
	}
	return ""
}

func (r *RadioButton) Focus() {
	r.focused = true
}

func (r *RadioButton) Blur() {
	r.focused = false
}

// GetSelectedIndex returns the index of the currently selected option
func (r *RadioButton) GetSelectedIndex() int {
	return r.selected
}

// SetSelected sets the selected option by index
func (r *RadioButton) SetSelected(index int) {
	if index >= 0 && index < len(r.options) {
		r.selected = index
		r.cursor = index
	}
}
