package commands

import (
	"os"
	"os/exec"
	"syscall"
)

func Exec(instructions []string) error {
	cmd, err := exec.LookPath(instructions[0])
	if err != nil {
		return err
	}
	if err = syscall.Exec(cmd, instructions[1:], os.Environ()); err != nil {
		return err
	}
	return nil
}
