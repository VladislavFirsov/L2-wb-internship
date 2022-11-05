package commands

import (
	"errors"
	"os"
	"path/filepath"
)

func changeDir(instructions []string) error {
	var err error
	var path string
	path, err = Pwd()
	if instructions[0] == ".." {
		path = filepath.Dir(path)
		if err = os.Chdir(path); err != nil {
			return err
		}
		return nil
	} else if err = os.Chdir(instructions[0]); err != nil {
		return errors.New("no such file or directory")
	}
	return nil
}
