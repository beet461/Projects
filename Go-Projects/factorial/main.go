package main

import "fmt"

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}
func main() {
	var i int
	fmt.Println("\n", "Type in your number.")
	fmt.Scan(&i)
	fmt.Printf("Factorial of %d is %d", i, factorial(i))
	main()
}
