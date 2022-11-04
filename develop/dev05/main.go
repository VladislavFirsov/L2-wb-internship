package main

import (
	"errors"
	"flag"
	"log"
)

var (
	ACapFlag int
	BFlag    int
	CFlag    int
	cntFlag  bool
	iFlag    bool
	vFlag    bool
	nFlag    bool
)

const path = "text.txt"

var ErrOutOfRange = errors.New("no more strings there")

func init() {
	flag.CommandLine.IntVar(&ACapFlag, "A", 0, "prints number of lines after the match")
	flag.CommandLine.IntVar(&BFlag, "B", 0, "prints number of lines before the match")
	flag.CommandLine.IntVar(&CFlag, "C", 0, "prints number of lines before and after the match")
	flag.CommandLine.BoolVar(&cntFlag, "c", false, "prints only a count of the lines that match a pattern")
	flag.CommandLine.BoolVar(&iFlag, "i", false, "ignores case for matches")
	flag.CommandLine.BoolVar(&vFlag, "v", false, "prints all the lines which don't match the pattern")
	flag.CommandLine.BoolVar(&nFlag, "n", false, "prints lines that matches and their numbers")
}

func main() {
	flag.Parse()
	pattern := flag.Arg(0)
	data := getDataFromFile(path, pattern)
	switch {
	case cntFlag:
		data.countMatches()
	case ACapFlag != 0:
		if err := data.handleAFlag(ACapFlag); err != nil {
			log.Fatalln(err)
		}
	case BFlag != 0:
		if err := data.handleBFlag(BFlag); err != nil {
			log.Fatalln(err)
		}
	case CFlag != 0:
		err1 := data.handleAFlag(CFlag)
		err2 := data.handleBFlag(CFlag)
		if err1 != nil || err2 != nil {
			log.Fatalln("error parsing CFlag")
		}
	}
}
