package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func getDataFromFile(path string) [][]string {
	data := make([][]string, 0)

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		data = append(data, s)
	}

	if scanner.Err() != nil {
		log.Fatalf("error from reading file: %s", err)
	}
	return data
}

func makeSortedFile(filename string, data [][]string) { //creates a file and writes to it
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	if uFlag {
		checker := make(map[string]int)
		for i := range data {
			str := strings.Join(data[i], " ")
			if checker[str] >= 1 {
				continue
			} else {
				checker[str]++
				file.WriteString(str + "\n")
			}

		}
		return
	}

	for i := range data {
		str := strings.Join(data[i], " ")
		file.WriteString(str + "\n")
	}
}
