package main

import "fmt"

type coffeeMaker interface {
	makeCoffee()
}

type coffeeMachine struct {
	state coffeeMaker
}

func (cm *coffeeMachine) makeCoffee() {
	cm.state.makeCoffee()
}

func (cm *coffeeMachine) setState(value coffeeMaker) {
	cm.state = value
}
func newCoffeeMachine() *coffeeMachine {
	return &coffeeMachine{state: &fullMachine{}}
}

type fullMachine struct {
}

func (full *fullMachine) makeCoffee() {
	fmt.Println("One coffee for you")
}

type emptyMachine struct {
}

func (empty *emptyMachine) makeCoffee() {
	fmt.Println("Sorry, the machine is empty")
}

func main() {
	coffeemachine1 := newCoffeeMachine()
	coffeemachine1.makeCoffee()
	coffeemachine1.setState(&emptyMachine{})
	coffeemachine1.makeCoffee()
}
