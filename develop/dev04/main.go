package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"пятак", "листок", "столик", "пятка", "тяпка", "слиток"}
	fmt.Println(GetAnagram(words))

}

func GetAnagram(words []string) map[string][]string {
	anagrams := make(map[string][]string, len(words))

	for _, value := range words {
		value = strings.ToLower(value)
	}

	for _, value := range words {

		isExists := false

		for key := range anagrams {
			if CompareStrings(key, value) {
				anagrams[key] = append(anagrams[key], value)
				isExists = true
				break
			}
		}
		if !isExists {
			anagrams[value] = append(anagrams[value], value)
		}
	}
	return anagrams
}

func GetLetters(word string) map[rune]int {
	letters := make(map[rune]int)
	for _, val := range word {
		letters[val]++
	}
	return letters
}

func CompareStrings(stringOne, stringTwo string) bool {
	letters := GetLetters(stringOne)
	for _, letter := range stringTwo {
		val, ok := letters[letter]
		if ok && val > 0 {
			letters[letter] -= 1
		}

	}
	for _, val := range letters {
		if val != 0 {
			return false
		}

	}
	return true
}
