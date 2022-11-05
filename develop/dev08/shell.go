package main

import (
	"bufio"
	"fmt"
	"github.com/VladislavFirsov/L2-wb-internship/commands"
	"io"
	"os/user"
	"strings"
)

type shell struct {
	reader io.Reader
	writer io.Writer
}

func newShell(reader io.Reader, writer io.Writer) *shell {
	return &shell{
		reader: reader,
		writer: writer,
	}
}

func (s *shell) start() error {
	user, err := user.Current()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(s.reader)

	for {
		fmt.Fprint(s.writer, user.Username+" ") // prints prefix

		scanner.Scan()
		str := scanner.Text() // getting input

		cmds := strings.Split(str, " ")

		err = commands.DoCommand(cmds[0], cmds[1:])
		if err != nil {
			return err
		}
	}
}
