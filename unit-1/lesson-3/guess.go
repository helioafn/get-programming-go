package main

import (
	"fmt"
	"math/rand"
)

// Write a guess-the-number program. Make the compuer pick random numbers between 1-100
// until it guesses your number, which you declare at the top of the program. Display each guess
// and whether it was too big or too small.
func main() {
	var myChoice = 71

	for {
		var computerGuess = rand.Intn(100) + 1
		if computerGuess == myChoice {
			fmt.Printf("You guessed it right! My guess was %v\n", myChoice)
			break
		} else if computerGuess > myChoice {
			fmt.Println("Your guess is too big, try a smaller number.")
		} else if computerGuess < myChoice {
			fmt.Println("Your guess is too small, try a bigger number.")
		}
	}
}
