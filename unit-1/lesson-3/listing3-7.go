package main

import "fmt"

func main() {
	var room = "lake"

	switch {
	case room == "lake":
		fmt.Println("You find yourself in a dimly lit cavern.")
		fallthrough // This keyword executes the body of the next 'case'
	case room == "underwater":
		fmt.Println("The water is freezing cold.")
	}

}
