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
		fmt.Println("You typed something in wrong.")
		main()
	}
	var oops float64
	oops = 404
	return oops
}

func main() {
	var function string
	var exit string
	var onum float64
	var senum float64
	fmt.Println("Type go to use the calculator or anything else to exit.")
	fmt.Scan(&exit)
	if exit != "go" {
		os.Exit(707)
	}
	fmt.Println("Choose a function, +, -, *, /, % (% is going to be the first number that you typed)")
	fmt.Scan(&function)
	fmt.Println("Choose the first number")
	fmt.Scan(&onum)
	fmt.Println("Choose the second number")
	fmt.Scan(&senum)
	start(function, onum, senum)
	main()
}
