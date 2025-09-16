package dto

import "github.com/daffadon/fndn/internal/types"

type Step struct {
	Label    string
	Input    types.Input
	Validate func(value any) error
}
