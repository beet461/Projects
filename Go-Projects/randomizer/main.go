package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func randomize(limid int) int {
	rand.Seed(time.Now().UnixNano())
	var word string
	var randomlimit int
	randomlimit = rand.Intn(limid)
	fmt.Println("Type in anything to begin! Or exit to exit. Or restart to restart")
	fmt.Scan(&word)
	if word == "exit" {
		os.Exit(404)
	} else if word == "restart" {
		main()
	}
	fmt.Println("\r Your number is", randomlimit)
	randomize(limid)
	return randomlimit
}

func main() {
	var limit int
	fmt.Println("Welcome to the randomizer! This is the place for random numbers.")
	fmt.Println("Choose your limit!")
	fmt.Scan(&limit)
	randomize(limit)
}
