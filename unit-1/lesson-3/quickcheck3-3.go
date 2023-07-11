package main

import "fmt"

// Write a program that uses if-else to display a description for each of three rooms: cave, entrance and mountain.

func main() {
	var choice = "mountain"

	if choice == "mountain" {
		fmt.Println("A mountain with an amazing view of the castle and plains.")
	} else if choice == "entrance" {
		fmt.Println("You can see the castle from quite far, it's getting more beautiful when you're getting close.")
	} else if choice == "cave" {
		fmt.Println("A shady cave, you wonder what's inside of it.")
	} else {
		fmt.Println("Didn't get that correctly.")
	}
}
