package commands

import (
	"fmt"
	"strings"
)

func Echo(instructions []string) {
	echo := strings.Join(instructions, " ")
	fmt.Println(echo)
}
