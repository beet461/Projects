package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func randomize(limid int) int {
	rand.Seed(time.Now().UnixNano())
	var randomlimit int
	switch limid {
	case 10:
		randomlimit = rand.Intn(10)
	case 100:
		randomlimit = rand.Intn(100)
	case 1000:
		randomlimit = rand.Intn(1000)
	case 10000:
		randomlimit = rand.Intn(10000)
	case 100000:
		randomlimit = rand.Intn(100000)
	case 1000000:
		randomlimit = rand.Intn(1000000)
	}
	return randomlimit
}

func mainbit(limig int) {
	j := randomize(limig)
	var word string
	fmt.Println("Type in anything to begin! Or exit to exit. Or restart to restart")
	fmt.Scan(&word)
	if word == "exit" {
		os.Exit(404)
	} else if word == "restart" {
		main()
	}
	fmt.Println("Your number is", j)
	mainbit(limig)
}

func main() {
	var limit int
	fmt.Println("Welcome to the randomizer! This is the place for random numbers.")
	fmt.Println("Choose your limit! 10, 100, 1000, 10000, 100000 or 1000000")
	fmt.Scan(&limit)
	mainbit(limit)
}
