package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/daffadon/fndn/internal/app"
	"github.com/daffadon/fndn/internal/infra"
	"github.com/daffadon/fndn/internal/ui"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Start interactive project initialization",
	RunE: func(cmd *cobra.Command, args []string) error {
		uc := &app.InitProjectUseCase{Runner: infra.ShellRunner{}}
		p := tea.NewProgram(ui.NewModel(uc))
		_, err := p.Run()
		return err
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
