package domain

import (
	"log"
	"os"

	"github.com/daffadon/fndn/internal/infra"
	template "github.com/daffadon/fndn/internal/template/framework"
	"golang.org/x/tools/imports"
)

func InitGin(i infra.CommandRunner, path *string) error {
	if path != nil {
		// Create directory
		i.Run("mkdir", []string{"-p", *path + "/config/router"}, "")

		// Define the Go file name
		fileName := *path + "/config/router/http.go"

		// Touch the file
		i.Run("touch", []string{fileName}, "")

		// Format the template with goimports
		opts := &imports.Options{
			Comments:  true,
			TabWidth:  8,
			TabIndent: true,
			Fragment:  false,
		}

		formatted, err := imports.Process(fileName, []byte(template.GinConfigTemplate), opts)
		if err != nil {
			log.Fatal(err)
		}

		// Write formatted content
		err = os.WriteFile(fileName, formatted, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
