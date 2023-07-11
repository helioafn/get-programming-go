package main

import "fmt"

func main() {
	fmt.Println("There is a sign near the entrance that reads 'No minors'.")
	var age = 25
	var isMinor = age < 18
	fmt.Printf("At age %v, am I a minor? %v\n", age, isMinor)
}
