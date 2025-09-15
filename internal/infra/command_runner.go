package infra

import "os/exec"

type CommandRunner interface {
	Run(name string, args ...string) error
}

type ShellRunner struct{}

func (ShellRunner) Run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}
