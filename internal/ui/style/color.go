package style

import "github.com/charmbracelet/lipgloss"

var (
	XStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Bold(true)
	StepLabelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Bold(true)
	ArrowStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)
	ErrorStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true)
)
