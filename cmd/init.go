package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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
	rootCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(frameworkCmd)
	generateCmd.AddCommand(dbCmd)
	generateCmd.AddCommand(mqCmd)
	generateCmd.AddCommand(cacheCmd)
	generateCmd.AddCommand(storageCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
