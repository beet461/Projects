package main

import (
	"fmt"
	"os"
)

func start(function string, onum float64, senum float64) float64 {
	switch function {
	case "+":
		fmt.Println("Your answer is", onum+senum)
	case "-":
		fmt.Println("Your answer is", onum-senum)
	case "*":
		fmt.Println("Your answer is", onum*senum)
	case "/":
		fmt.Println("Your answer is", onum/senum)
	case "%":
		fmt.Println("Your answer is", (onum/100)*senum)
	default:
		os.Exit(707)
	}
	var oops float64
	oops = 404
	return oops
}

func main() {
	var function string
	var onum float64
	var senum float64
	fmt.Println("Choose a function, +, -, *,/, % (% is going to be the first number that you typed) or anything else to exit.")
	fmt.Scan(&function)
	fmt.Println("Choose the first number")
	fmt.Scan(&onum)
	fmt.Println("Choose the second number")
	fmt.Scan(&senum)
	start(function, onum, senum)
	main()
}
