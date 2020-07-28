package main

import "fmt"


func frac(funx string, nu3, nu4 float64) float64{
	switch funx {
	case "+":
		fmt.Println(nu3 + nu4)
	case "-":
		fmt.Println(nu3 - nu4)
	case "*":
		fmt.Println(nu3 * nu4)
	case "/":
		fmt.Println(nu3 / nu4)
}

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
		fmt.Println((nu1 / 100) * nu2))
	case "//":
		fmt.Println("Choose your second fraction.")
		nu3 := nu1 + nu2
		fmt.Scan(&nu1, &nu2)
		nu4 := nu1 + nu2
		fmt.Println("And now choose your function (no, you cant use // or %).")
		fmt.Scan(&funx)
	}
}
