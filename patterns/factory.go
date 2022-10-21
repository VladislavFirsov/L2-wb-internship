package main

import "fmt"

type transporter interface {
	deliver()
}

type transport struct {
	name string
}

func (t *transport) deliver() {
	fmt.Printf("delivered by %s\n", t.name)
}

type scooter struct {
	transport
}

func newScooter() transporter {
	return &scooter{
		transport{name: "scooter"},
	}
}

type bicycle struct {
	transport
}

func newBicycle() transporter {
	return &bicycle{
		transport{name: "bicycle"},
	}
}

func makeTransport(name string) (transporter, error) {
	switch name {
	case "scooter":
		return newScooter(), nil
	case "bicycle":
		return newBicycle(), nil
	default:
		return nil, fmt.Errorf("something went wrong")
	}
}

func main() {
	bicycle, err := makeTransport("bicycle")
	scooter, _ := makeTransport("scooter")
	if err != nil {
		fmt.Println(err)
	}
	bicycle.deliver()
	scooter.deliver()
}
