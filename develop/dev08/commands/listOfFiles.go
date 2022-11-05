package commands

import (
	"io/ioutil"
	"strings"
)

func Ls() (string, error) {
	var builder strings.Builder
	path, err := Pwd()
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return "", err
	}
	for _, file := range files {
		builder.WriteString(file.Name() + "\n")
	}
	return builder.String(), nil
}
