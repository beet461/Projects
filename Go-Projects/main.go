package main

import "fmt"

func main() {
	var funx string
	var nu1 float64
	var nu2 float64
	fmt.Println("Choose your function, +, -, *, /, %, //(fraction) or exit")
	fmt.Scan(&funx)
	fmt.Println("Choose your numbers")
	fmt.Scan(&nu1)
	fmt.Print(funx)
	fmt.Scan(&nu2)
	fmt.Println(nu1, nu2)
}
