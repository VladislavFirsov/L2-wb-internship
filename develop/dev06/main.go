package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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

func getDataFromStdin() []string {
	data := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return data
}

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
