package commands

import (
	"os"
	"strconv"
)

func killProcess(instructions []string) error {
	var err error
	pid, err := strconv.Atoi(instructions[0])
	if err != nil {
		return err
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	err = process.Kill()
	if err != nil {
		return err
	}
	return nil
}
