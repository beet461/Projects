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
	fmt.Scan(&frac12)
	fmt.Println("Choose the second fraction.")
	fmt.Scan(&frac21)
	fmt.Scan(&frac22)
	fracans := (frac11 / frac12) / (frac21 / frac22)
	fracadd := (frac11 / frac12) + (frac21 / frac22)
	fracsub := (frac11 / frac12) - (frac21 / frac22)
	fracmul := (frac11 / frac12) * (frac21 / frac22)
	switch typ {
	case "-+":
		fmt.Println(fracadd)
	case "--":
		fmt.Println(fracsub)
	case "-*":
		fmt.Println(fracmul)
	case "-/":
		fmt.Println(fracans)
	}
	return 404
}

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
	fmt.Println("Choose a function, +, -, *, /, % (% is going to be the first number that you typed) or - (for fraction and put the function you want after it)")
	fmt.Scan(&function)
	if function == "-" {
		fraction(function)
	}
	fmt.Println("Choose the first number")
	fmt.Scan(&onum)
	fmt.Println("Choose the second number")
	fmt.Scan(&senum)
	start(function, onum, senum)
	main()
}
