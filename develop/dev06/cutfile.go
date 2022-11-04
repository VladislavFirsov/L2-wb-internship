package main

import (
	"fmt"
	"strings"
)

func cutLine(data string, delimiter string) {
	if delimiter == "" || !strings.Contains(data, delimiter) {
		if !sFlag {
			fmt.Println(data)
		}
		return
	}

	fields := strings.Split(data, delimiter)

	if fFlag > len(fields) {
		fmt.Println()
	} else {
		fmt.Println(fields[fFlag-1])
	}
}
