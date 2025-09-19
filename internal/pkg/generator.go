package pkg

import (
	"os"
	"strings"

	"github.com/daffadon/fndn/internal/infra"
	"golang.org/x/tools/imports"
)

func GoFileGenerator(i infra.CommandRunner,
	path *string,
	folderName,
	fileName,
	template string) error {
	// Ensure parent directory exists
	dir := *path + folderName
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Define the Go file name
	fn := *path + fileName

	// Create the file if it doesn't exist (Go way)
	f, err := os.OpenFile(fn, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	f.Close()

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

func GenericFileGenerator(i infra.CommandRunner,
	path *string,
	folderName,
	fileName,
	template string) error {

	// Define the YAML file name
	fn := *path + fileName

	// Ensure parent directory exists
	dir := *path + folderName
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Touch the file
	f, err := os.OpenFile(fn, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	f.Close()

	cleanTemplate := strings.ReplaceAll(template, "\t", "  ")

	// Write the sanitized YAML template directly to the file
	err = os.WriteFile(fn, []byte(cleanTemplate), 0644)
	if err != nil {
		return err
	}
	return nil
}
