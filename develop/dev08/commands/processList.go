package commands

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"strings"
)

func ShowProcessList() (string, error) {
	var builder strings.Builder
	processes, err := ps.Processes()
	if err != nil {
		return "", err
	}
	for _, proc := range processes {
		builder.WriteString(fmt.Sprintf("%d %s\n", proc.Pid(), proc.Executable()))
	}
	return builder.String(), nil
}
