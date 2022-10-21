package main

import "fmt"

type laptop struct {
}

func (l *laptop) loadOs() {
	fmt.Println("Turning on...")
}

func (l *laptop) loadScreen() {
	fmt.Println("getting screen ready...")
}

func (l *laptop) terminateAll() {
	fmt.Println("turning off...")
}

type laptopFacade struct {
	laptop
}

func (l *laptopFacade) turnOn() {
	l.laptop.loadOs()
	l.laptop.loadScreen()
}

func (l *laptopFacade) turnoff() {
	l.laptop.terminateAll()
}

func newLaptop() *laptopFacade {
	return &laptopFacade{}
}

func main() {
	laptop := newLaptop()
	laptop.turnOn()
	laptop.turnoff()
}
