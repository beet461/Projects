package main

import (
	"fmt"
	"os"
)

func fraction(typ string) float64 {
	var frac11 float64
	var frac12 float64
	var frac21 float64
	var frac22 float64
	fmt.Println("Choose the first fraction.")
	fmt.Scan(&frac11)
	fmt.Println("/")
	fmt.Scan(&frac12)
	fmt.Println("Choose the second fraction.")
	fmt.Scan(&frac21)
	fmt.Println("/")
	fmt.Scan(&frac22)
	switch typ {
	case "-+":
		fmt.Println((frac11 / frac12) + (frac21 / frac22))
	case "--":
		fmt.Println((frac11 / frac12) - (frac21 / frac22))
	case "-*":
		fmt.Println((frac11 / frac12) * (frac21 / frac22))
	case "-/":
		fmt.Println((frac11 / frac12) / (frac21 / frac22))
	}
	return 404
}

func start(function string, firnum float64, secnum float64) float64 {
	switch function {
	case "+":
		fmt.Println("Your answer is", firnum+secnum)
	case "-":
		fmt.Println("Your answer is", firnum-secnum)
	case "*":
		fmt.Println("Your answer is", firnum*secnum)
	case "/":
		fmt.Println("Your answer is", firnum/secnum)
	case "%":
		fmt.Println("Your answer is", (firnum/100)*secnum)
	default:
		fmt.Println("You typed something in wrong.")
		main()
	}
	var oops float64
	oops = 404
	return oops
}

func hi() {
	fmt.Println("HI")
}

func main() {
	var function string
	var exit string
	var firnum float64
	var secnum float64
	fmt.Println("Type go to use the calculator or anything else to exit.")
	fmt.Scan(&exit)
	if exit != "go" {
		os.Exit(707)
	}
	fmt.Println("Choose a function, +, -, *, /, % (% is going to be the first number that you typed) or - (for fraction and put the function you want after it)")
	fmt.Scan(&function)
	if function == "-+" || function == "--" || function == "-*" || function == "-/" {
		fraction(function)
	} else if function == "-" {
		fmt.Println("Not a valid function.")
		main()
	} else {
		fmt.Println("Choose the first number")
		fmt.Scan(&firnum)
		fmt.Println("Choose the second number")
		fmt.Scan(&secnum)
		start(function, firnum, secnum)
	}
	main()
}
