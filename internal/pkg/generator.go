package pkg

import (
	"os"

	"github.com/daffadon/fndn/internal/infra"
	"golang.org/x/tools/imports"
)

func FileGenerator(i infra.CommandRunner,
	path *string,
	folderName,
	fileName,
	template string) error {
	// Create directory
	i.Run("mkdir", []string{"-p", *path + folderName}, "")

	// Define the Go file name
	fn := *path + fileName

	// Touch the file
	i.Run("touch", []string{fn}, "")

	opts := &imports.Options{
		Comments:  true,
		TabWidth:  8,
		TabIndent: true,
		Fragment:  false,
	}

	formatted, err := imports.Process(fn, []byte(template), opts)
	if err != nil {
		return err
	}

	err = os.WriteFile(fn, formatted, 0644)
	if err != nil {
		return err
	}
	return nil
}
