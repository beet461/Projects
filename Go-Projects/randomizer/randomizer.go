package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var word string
	randomlimit := rand.Intn(100)
	fmt.Println("Welcome to the randomizer! This is the place for random numbers.")
	fmt.Println("Type in randomize to begin!")
	fmt.Scan(&word)
	for word != "randomize" {
		fmt.Println("Try again")
		fmt.Scan(&word)
	}
	fmt.Println("Your number is", randomlimit)
}
