package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	err := validateWget(args[1])
	if err != nil {
		log.Fatalln(err)
	}
	response, err := getWebSite(args[2])
	if err != nil {
		log.Fatalln(err)
	}
	size, err := writeToFile(response, "downloaded webpage")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Downloaded size is:", size)
}
