package main

import (
	"fmt"
	"math/rand" // This is the import path for the rand package.
)

// Random numbers
func main() {
	// rand.Intn(n) -> returns a number from [0 to n-1]
	var num = rand.Intn(10) + 1
	fmt.Println(num)
	num = rand.Intn(10) + 1
	fmt.Println(num)
}
