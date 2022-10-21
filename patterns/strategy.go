package main

import (
	"fmt"
	"math/rand"
	"time"
)

type strategyI interface {
	sort([]int)
}
type bubbleSort struct {
}

func (b *bubbleSort) sort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}

type quicksort struct {
}

func (q *quicksort) sort(numbers []int) {
	if len(numbers) <= 1 {
		return
	}

	start, end := 0, len(numbers)-1
	pivot := end
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < numbers[pivot] {
			numbers[start], numbers[i] = numbers[i], numbers[start]
			start++
		}
	}
	numbers[start], numbers[pivot] = numbers[pivot], numbers[start]

	q.sort(numbers[:start])
	q.sort(numbers[start+1:])
}

type strategy struct {
	algorithm strategyI
}

func main() {
	numbers1 := makeSlice()
	time.Sleep(1 * time.Millisecond) // to make both slices random
	numbers2 := makeSlice()
	s := strategy{algorithm: &bubbleSort{}}
	s.algorithm.sort(numbers1)
	fmt.Println("Bublesorted", numbers1)
	s.algorithm = &quicksort{}
	s.algorithm.sort(numbers2)
	fmt.Println("Quicksorted", numbers2)

}

func makeSlice() []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, 10)
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(100)
	}
	return slice
}
