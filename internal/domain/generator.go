package domain

import (
	"fmt"

	"github.com/daffadon/fndn/internal/pkg"
)

type GeneratorType string

const (
	GeneratorFramework GeneratorType = "framework"
	GeneratorDatabase  GeneratorType = "database"
	GeneratorMQ        GeneratorType = "mq"
	GeneratorCache     GeneratorType = "cache"
	GeneratorStorage   GeneratorType = "storage"
)

type Generator struct {
	Type  GeneratorType
	Value string
}

// lookup returns the template for value or a loud error on unknown input.
// Empty templates used to be written silently; now they fail fast.
func lookup(table map[string]string, tech, value string) (string, error) {
	t, ok := table[value]
	if !ok {
		return "", fmt.Errorf("unknown %s %q", tech, value)
	}
	return t, nil
}

// lookupFile is lookup for [filename template] pairs. Unknown values used to
// write empty files via the || guard; now they fail fast.
func lookupFile(table map[string][2]string, tech, value string) ([2]string, error) {
	file, ok := table[value]
	if !ok {
		return [2]string{}, fmt.Errorf("unknown %s %q", tech, value)
	}
	return file, nil
}

// resolveTarget handles GenerateSpecific* filename collision: if the default
// file exists, suffix with the tech value instead of overwriting.
func resolveTarget(defaultFile, prefix, value string) string {
	if pkg.IsFileExists("." + defaultFile) {
		return prefix + "_" + value + ".go"
	}
	return defaultFile
}
