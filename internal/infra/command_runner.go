package infra

import (
	"bytes"
	"fmt"
	"os/exec"
)

type CommandRunner interface {
	Run(name string, args []string, dir string) error
}
type runner struct{}

func NewCommandRunner() CommandRunner {
	return &runner{}
}

func (r *runner) Run(name string, args []string, dir string) error {
	cmd := exec.Command(name, args...)
	if dir != "" {
		cmd.Dir = dir
	}
	var errBuf bytes.Buffer
	cmd.Stderr = &errBuf

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("project generation failed \n\n%s", errBuf.String())
	}
	return nil
}
