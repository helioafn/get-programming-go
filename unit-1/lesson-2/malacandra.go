package main

import "fmt"

// Write a program to determine how fast a ship would need to travel (in km/h) in order to reach Malacandra in 28 days
// Assume a distance of 56,000,000 km.
func main() {
	const distance = 56000000 // km
	daysInTravel := 28
	hoursInADay := 24
	travelSpeed := distance / (daysInTravel * hoursInADay)
	fmt.Printf("Necessary speed to arrive Malacandra in %v days is: %v km/h\n", daysInTravel, travelSpeed)
}
