package main

import (
	"fmt"
	"math/rand"
)

// Implement a countdown where every second there's a 1 in 100 chance that the launch fails and the countdown stops.
func main() {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		count--
		if rand.Intn(100)+1 == 1 {
			break
		}
	}
	if count == 0 {
		fmt.Println("Liftoff!")
	} else {
		fmt.Println("Launch failed.")
	}
}
