package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func randomize() string {
	rand.Seed(time.Now().UnixNano())
	randomlimit := rand.Intn(100)
	k := fmt.Sprintln("Your number is", randomlimit)
	return k
}

func main() {
	j := randomize()
	var word string
	r := "Type in go for another number or anything else to exit"
	fmt.Println(r)
	fmt.Scan(&word)
	if word != "go" {
		os.Exit(404)
	} else {
		fmt.Println(j)
		main()
	}
}
