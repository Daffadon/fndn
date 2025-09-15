package pkg

import (
	"strings"
)

func LastSegment(module string) string {
	parts := strings.Split(module, "/")
	return parts[len(parts)-1]
}