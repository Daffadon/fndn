package pkg

import (
	"os"
	"strings"

	"golang.org/x/tools/imports"
)

func IsFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// ensureFile creates parent dir and touches file, returning full path.
func ensureFile(path *string, folderName, fileName string) (string, error) {
	dir := *path + folderName
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	fn := *path + fileName
	f, err := os.OpenFile(fn, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", err
	}
	f.Close()
	return fn, nil
}

func GoFileGenerator(
	path *string,
	folderName,
	fileName,
	template string) error {
	fn, err := ensureFile(path, folderName, fileName)
	if err != nil {
		return err
	}

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

func GenericFileGenerator(
	path *string,
	folderName,
	fileName,
	template string) error {
	fn, err := ensureFile(path, folderName, fileName)
	if err != nil {
		return err
	}

	cleanTemplate := template
	if strings.HasSuffix(fileName, ".yaml") || strings.HasSuffix(fileName, ".yml") {
		cleanTemplate = strings.ReplaceAll(template, "\t", "  ")
	}

	// Write the sanitized YAML template directly to the file
	err = os.WriteFile(fn, []byte(cleanTemplate), 0644)
	if err != nil {
		return err
	}
	return nil
}
