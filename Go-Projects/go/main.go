package main

import "fmt"

type onstruct struct {
	hi string
}

func main() {
	var yr onstruct
	fmt.Scan(&yr)
	yr.hi = yr.hi
	fmt.Println("you wrote %v", yr.hi)

}
