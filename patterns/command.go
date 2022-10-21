package main

import "fmt"

type command interface {
	execute()
}

type restaurant struct {
	totalDishes   int
	cleanedDishes int
}

func newRestaurant() *restaurant {
	return &restaurant{
		totalDishes:   10,
		cleanedDishes: 10,
	}
}

func (r *restaurant) makePizza(n int) command {
	return &makePizzaCommand{
		number:     n,
		restaurant: r,
	}
}

func (r *restaurant) makeSalad(n int) command {
	return &makeSaladCommand{
		number:     n,
		restaurant: r,
	}
}

func (r *restaurant) cleanDishes() command {
	return &cleanDishesCommand{
		restaurant: r,
	}
}

type makePizzaCommand struct {
	number     int
	restaurant *restaurant
}

func (pizza *makePizzaCommand) execute() {
	pizza.restaurant.cleanedDishes -= pizza.number
	fmt.Println("Made one pizza")

}

type makeSaladCommand struct {
	number     int
	restaurant *restaurant
}

func (salad *makeSaladCommand) execute() {
	salad.restaurant.cleanedDishes -= salad.number
	fmt.Println("Made one salad")
}

type cleanDishesCommand struct {
	restaurant *restaurant
}

func (c *cleanDishesCommand) execute() {
	c.restaurant.cleanedDishes = c.restaurant.totalDishes
	fmt.Println("Dishes cleaned")
}

func main() {
	clodMone := newRestaurant()
	for i := 0; i < 3; i++ {
		clodMone.makePizza(1).execute()
		clodMone.makeSalad(1).execute()
		fmt.Println(clodMone.cleanedDishes)
		clodMone.cleanDishes().execute()
		fmt.Println(clodMone.cleanedDishes)
	}
}
