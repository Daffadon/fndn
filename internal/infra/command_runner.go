package infra

import "os/exec"

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
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}
