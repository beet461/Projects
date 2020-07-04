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

func main() {
	var word string
	j := randomize()
	fmt.Println("Welcome to the randomizer! This is the place for random numbers.")
	fmt.Println("Type in anything to begin!")
	fmt.Scan(&word)
	fmt.Println("Your number is", j)
	fmt.Println("Type in go for another number or anything else to exit")
	fmt.Scan(&word)
	if word == "go" {
		main()
	} else {
		os.Exit(404)
	}
}
