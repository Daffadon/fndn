package domain

import (
	"fmt"

	"github.com/daffadon/fndn/internal/pkg"
	"github.com/daffadon/fndn/internal/types"
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

// None deselects an entity. Init* funcs treat it as "skip, generate nothing".
const None = types.None

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

// section is lookup that yields "" for None so yaml/compose builders can
// omit the block instead of failing.
func section(table map[string]string, tech, value string) (string, error) {
	if value == None {
		return "", nil
	}
	return lookup(table, tech, value)
}

// sectionFile is lookupFile that yields zero value for None.
func sectionFile(table map[string][2]string, tech, value string) ([2]string, error) {
	if value == None {
		return [2]string{}, nil
	}
	return lookupFile(table, tech, value)
}

// resolveTarget handles GenerateSpecific* filename collision: if the default
// file exists, suffix with the tech value instead of overwriting.
func resolveTarget(defaultFile, prefix, value string) string {
	if pkg.IsFileExists("." + defaultFile) {
		return prefix + "_" + value + ".go"
	}
	return defaultFile
}
