package main

import "fmt"

func main() {
	var funx string
	var nu1 float64
	var nu2 float64
	fmt.Println("Choose your function, +, -, *, /, %, //(fraction) or exit.")
	fmt.Scan(&funx)
	fmt.Println("Choose your numbers (space needed).")
	fmt.Scan(&nu1, &nu2)
	switch funx {
	case "+":
		fmt.Println(nu1 + nu2)
	case "-":
		fmt.Println(nu1 - nu2)
	case "*":
		fmt.Println(nu1 * nu2)
	case "/":
		fmt.Println(nu1 / nu2)
	case "%":
		fmt.Println((nu1 / 100) * nu2)
	}
}
