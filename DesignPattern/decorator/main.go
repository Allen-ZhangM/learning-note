package main

import "fmt"

func main() {

	var pizza_ food
	pizza_ = &pizza{}

	//Add cheese topping
	pizzaWithCheese := &cheeseTopping{
		pizza: pizza_,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of pizza with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
