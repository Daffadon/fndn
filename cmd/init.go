package cmd

import (
	"fmt"
	"os"

	"github.com/daffadon/fndn/internal/ui"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [.]",
	Short: "Initialize a new Go project",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var targetPath string
		if len(args) == 1 && args[0] == "." {
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}
			targetPath = cwd
		} else {
			targetPath = ""
		}
		return ui.RunModuleInput(targetPath)
	},
}


var rootCmd = &cobra.Command{
	Use:   "fndn",
	Short: "fndn - Foundation for Go projects",
	Long: `fndn helps you bootstrap Go backend projects with clean architecture.
It provides commands to initialize modules, generate boilerplate, and scaffold features.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run `fndn --help` to see available commands.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
