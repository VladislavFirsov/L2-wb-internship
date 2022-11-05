package main

import (
	"log"
	"os"
)

func main() {
	myShell := newShell(os.Stdin, os.Stdout)
	if err := myShell.start(); err != nil {
		log.Println(err)
	}

}
