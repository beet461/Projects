package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func randomize() int {
	rand.Seed(time.Now().UnixNano())
	randomlimit := rand.Intn(100)
	return randomlimit
}

func mainbit() {
	j := randomize()
	var word string
	fmt.Println("Type in anything to begin! Or exit to exit")
	fmt.Scan(&word)
	if word == "exit" {
		os.Exit(404)
	}
	fmt.Println("Your number is", j)
	mainbit()
}

func main() {
	fmt.Println("Welcome to the randomizer! This is the place for random numbers.")
	mainbit()
}
