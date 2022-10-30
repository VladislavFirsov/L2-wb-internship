package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

var (
	kFlag int
	nFlag bool
	rFlag bool
	uFlag bool
)
var ErrIncorrectColumnNumber = errors.New("invalid number of column")
var ErrAbsentInteger = errors.New("no integers found")

const path = "text.txt"

func init() {
	flag.CommandLine.IntVar(&kFlag, "k", 0, "specify a column to sort")
	flag.CommandLine.BoolVar(&nFlag, "n", false, "sort by numeric value")
	flag.CommandLine.BoolVar(&rFlag, "r", false, "reverse sorting")
	flag.CommandLine.BoolVar(&uFlag, "u", false, "ignore duplicate strings")
}

func main() {
	flag.Parse()
	data := getDataFromFile(path)
	if err := sort(data); err != nil {
		log.Fatalln(err)
	}
	if rFlag {
		reverseSort(data)
	}
	fmt.Println(data)
	makeSortedFile("sorted.txt", data)

}
