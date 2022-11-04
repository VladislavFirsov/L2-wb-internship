package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	fFlag int
	dFlag string
	sFlag bool
)

func init() {
	flag.CommandLine.IntVar(&fFlag, "f", 0, "prints out the selected fields of a text")
	flag.CommandLine.StringVar(&dFlag, "d", "\t", "prints the fields specified by delimiter")
	flag.CommandLine.BoolVar(&sFlag, "s", false, "prints lines that contains delimiter")
}

func main() {
	flag.Parse()
	data := getDataFromStdin()

	if fFlag <= 0 {
		log.Fatalln("invalid number of the field")
	}
	fmt.Println(data)
	for _, val := range data {
		cutLine(val, dFlag)
	}
}
