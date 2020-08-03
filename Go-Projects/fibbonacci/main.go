package main

import "fmt"

func fibonacci(i int) (ret int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}
func main() {
	var i int
	var in int
	fmt.Println("Choose how many numbers of the fibonacci sequence you want to know.")
	fmt.Scan(&in)
	fmt.Print("____")
	for i = 0; i < in; i++ {
		fmt.Printf("%d ", fibonacci(i))
		fmt.Println(" ")
	}
}
