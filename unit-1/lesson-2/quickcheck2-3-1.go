package main

import "fmt"

func main() {
	var travelSpeed = 100800.0 // km/h
	const hoursPerDay = 24
	var distance = 96300000.0 // km
	fmt.Printf("It would take %v days to reach Mars", (distance/travelSpeed)/hoursPerDay)
}
