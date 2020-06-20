package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomize() string {
	rand.Seed(time.Now().UnixNano())
	randomlimit := rand.Intn(100)
	k := fmt.Sprintln("Your number is", randomlimit)
	fmt.Println(k)
	return k
}

func main() {
	j := randomize()
	var word string
	r := "Type in go for another number or anything else to exit"
	fmt.Println(r)
	fmt.Scan(&word)
	if word == "go"
	fmt.Println(j)
	main()

}
