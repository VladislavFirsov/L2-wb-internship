package main

import "strconv"

func sort(data [][]string) error {
	if kFlag != 0 {
		if kFlag > len(data[0]) {
			return ErrIncorrectColumnNumber
		}
		bubbleSort(kFlag, data)
		return nil
	}
	if nFlag {
		indx, err := findDigitIndex(data)
		if err != nil {
			return err
		}
		bubbleSort(indx, data)
		return nil
	}

	bubbleSort(1, data)
	return nil
}

func findDigitIndex(data [][]string) (int, error) {
	for i := range data {
		for j := range data[i] {
			if _, err := strconv.Atoi(data[i][j]); err == nil {
				return j + 1, nil
			}
		}
	}
	return -1, ErrAbsentInteger
}

func bubbleSort(flag int, data [][]string) {
	indx := flag - 1
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {
			if compare(data[i][indx], data[j][indx]) {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
}

func compare(str1, str2 string) bool {
	num1, err := strconv.Atoi(str1)
	if err == nil {
		num2, _ := strconv.Atoi(str2)
		return num1 < num2
	}
	b1 := byte(str1[0])
	b2 := byte(str2[0])
	if b1 < b2 {
		return true
	}
	return false
}

func reverseSort(data [][]string) {
	for i := 0; i < len(data)/2; i++ {
		j := len(data) - i - 1
		data[i], data[j] = data[j], data[i]
	}
}
