package main

import (
	"fmt"
	"math/rand"
)

// Distance between Earth and Mars varies from nearby to opposite sides of the sun.
// Write a program that generates a random distance from 56,000,000 to 401,000,000 km.
func main() {
	var distance = rand.Intn(345000001) + 56000000
	fmt.Println(distance)
}
