package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func getDataFromFile(filename string, pattern string) *linkedList {
	list := &linkedList{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		node := list.createNode(str)
		list.search(node, strings.Split(str, " "), pattern)
		list.insert(node)

	}
	if scanner.Err() != nil {
		log.Fatalln("error while reading file")
	}
	return list
}
