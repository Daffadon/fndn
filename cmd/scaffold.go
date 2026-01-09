package cmd

import (
	"os"

	"github.com/daffadon/fndn/internal/ui/cli/main_init"
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
		return main_init.RunModuleInput(targetPath)
	},
}
