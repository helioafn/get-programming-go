package main

import (
	"fmt"
	"strings" // Package that contains functions to operate with strings
)

// A function that returns a Boolean value

func main() {
	fmt.Println("You find yourself in a dimly lit cavern.")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")

	fmt.Println("You leave the cave:", exit)
}
