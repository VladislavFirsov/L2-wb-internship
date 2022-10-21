package main

import "fmt"

type namer interface {
	introduce()
	setNext(handler namer)
}

type fullName struct {
	name string
	next namer
}

func (fn *fullName) introduce() {
	fmt.Printf("%s ", fn.name)
	if fn.next != nil {
		fn.next.introduce()
	}
}

func (fn *fullName) setNext(next namer) {
	fn.next = next
}

func newFullName(name string) *fullName {
	return &fullName{
		name: name,
		next: nil,
	}
}

func main() {
	firstName := newFullName("mister")
	lastName := newFullName("Anderson")
	firstName.setNext(lastName)
	firstName.introduce()
}
