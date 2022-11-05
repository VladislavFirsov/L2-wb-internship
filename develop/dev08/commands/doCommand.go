package commands

import (
	"fmt"
	"os"
)

func DoCommand(command string, instructions []string) error {
	var data string
	var err error
	switch command {
	case "cd":
		err = changeDir(instructions)
	case "pwd":
		data, err = Pwd()
		fmt.Println(data)
	case "echo":
		Echo(instructions)
	case "kill":
		err = killProcess(instructions)
	case "ps":
		data, err = ShowProcessList()
		fmt.Println(data)
	case "exec":
		err = Exec(instructions)
	case "quit":
		fmt.Println("exit the shell")
		os.Exit(0)
	case "ls":
		data, err = Ls()
		fmt.Println(data)
	default:
		fmt.Println("Unknown command")
	}
	if err != nil {
		return err
	}
	return nil
}
