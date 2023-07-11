package main
import "fmt"
func main() {
	// Grouping multiple variables at once
	//var (
	//	distance = 56000000
	//	speed = 100800
	//)

	// Also you can declare multiple variables on a single line
//	var distance, speed = 56000000, 100800

	// Assignment operators
	var weight = 149.0
//	weight = weight * 0.3783
	weight *= 0.3783

	// Increment operator
	var age = 25
//	age = age + 1
//	age += 1
	age++

	// You can also x-- or price /= 2 in the same way as above
	// !! Go doesn't support ++count (thx god)
