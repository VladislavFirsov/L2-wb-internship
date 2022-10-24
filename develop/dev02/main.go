package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var ErrIncorrectString = errors.New("invalid string")

func main() {
	answer, err := UnpackString(`abcd\`)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(answer)
}

func UnpackString(s string) (string, error) {
	if len(s) == 0 { // if len string is 0
		return "", nil
	}

	stringOfRunes := []rune(s)
	esc := false // will keep track of escape sequence
	result := "" // final string

	if unicode.IsDigit(stringOfRunes[0]) { // if invalid string from the beginning
		return "", ErrIncorrectString
	}

	for i, value := range stringOfRunes {

		if esc { // add a value after escape sequence
			result += string(value)
			esc = false
			continue
		}
		if value == '\\' { // start of an escape sequence
			if unicode.IsLetter(stringOfRunes[i+1]) {
				return "", ErrIncorrectString
			}
			esc = true
			continue
		}

		if unicode.IsDigit(value) { // unpacking rune or escape letter
			num, err := strconv.Atoi(string(value))
			if err != nil {
				log.Println(err)
				return "", err
			}
			if num == 0 {
				result = result[:len(result)-1]
				continue
			}
			result += strings.Repeat(string(stringOfRunes[i-1]), num-1)
			continue

		}
		result += string(value) // adding a single letter
	}

	return result, nil
}
