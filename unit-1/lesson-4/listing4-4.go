package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Short declaration allows us to declare in a specific scope. Like an if statement or for loop.
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space Adventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	} // num is no longer in scope
}
