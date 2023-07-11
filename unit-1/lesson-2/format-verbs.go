package main

import "fmt"

func main() {
	// Positive values for padding, pads with spaces to the left, negative values pads with spaces to the right.
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}
