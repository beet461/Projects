package main

import (
	"fmt"
	"os"
)

func fraction(typ string) float64 {
	var fracs [4]float64
	fmt.Println("Choose the first fraction.")
	fmt.Scan(&fracs[0])
	fmt.Println("/")
	fmt.Scan(&fracs[1])
	fmt.Println("Choose the second fraction.")
	fmt.Scan(&fracs[2])
	fmt.Println("/")
	fmt.Scan(&fracs[3])
	switch typ {
	case "-+":
		fmt.Println((fracs[0] / fracs[1]) + (fracs[2] / fracs[3]))
	case "--":
		fmt.Println((fracs[0] / fracs[1]) - (fracs[2] / fracs[3]))
	case "-*":
		fmt.Println((fracs[0] / fracs[1]) * (fracs[2] / fracs[3]))
	case "-/":
		fmt.Println((fracs[0] / fracs[1]) / (fracs[2] / fracs[3]))
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
	fmt.Println("Choose a function, +, -, *, /, % (% is going to be the first number that you typed), - (for fraction and put the function you want after it) or mod(modulus)")
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
