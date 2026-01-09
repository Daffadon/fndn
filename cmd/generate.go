package cmd

import (
	"fmt"

	cli_database "github.com/daffadon/fndn/internal/ui/cli/main_generate/database"
	cli_framework "github.com/daffadon/fndn/internal/ui/cli/main_generate/framework"
	cli_mq "github.com/daffadon/fndn/internal/ui/cli/main_generate/mq"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate tools config for framework, database, message queue, cache, or storage",
}

var frameworkCmd = &cobra.Command{
	Use:   "framework",
	Short: "Generate a go http",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli_framework.RunGenerateFrameworkConfig()
	},
}

var dbCmd = &cobra.Command{
	Use:   "database",
	Short: "Generate a database",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli_database.RunGenerateDatabaseConfig()
	},
}

var mqCmd = &cobra.Command{
	Use:   "mq",
	Short: "Generate a message queue",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli_mq.RunGenerateMQConfig()
	},
}

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Generate a cache",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Add logic to generate a cache
		fmt.Println("Generating cache...")
		return nil
	},
}

var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Generate a file storage",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Add logic to generate file storage
		fmt.Println("Generating file storage...")
		return nil
	},
}
