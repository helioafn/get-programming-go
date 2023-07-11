package main

import "fmt"

func main() {
	// Concise form of switch
	var room = "lake"

	switch room {
	case "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case "underwater":
		fmt.Println("The water is freezing cold.")
	default:
		fmt.Println("Didn't get that quite right.")
	}

}
