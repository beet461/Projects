package main

import "fmt"

func add(a, b float64) float64 {
	return a +b
}

func sub(a, b float64) float64 {
	return a - b
}

func times(a, b float64) float64 {
	return a * b
}

func divi(a, b float64) float64 {
	return a / b
}

func main() {
	var choice string
	var firstnum float64
	var secondnum float64
	onenum := "Choose the first number"
	secnum := "Choose the second number"
	fmt.Println("Welcome to the calculator!")
	fmt.Println("Choose a function, +, -, * or /")
	fmt.Scan(&choice)
	if choice != "" {
		switch choice {
		case "+":
			fmt.Println(onenum)
			fmt.Scan(&firstnum)
			fmt.Println(secnum)
			fmt.Scan(&secondnum)
			j := add(firstnum, secondnum)
			fmt.Println("Your answer is ", j)
		case "-":
			fmt.Println(onenum)
			fmt.Scan(&firstnum)
			fmt.Println(secnum)
			fmt.Scan(&secondnum)
			j := sub(firstnum, secondnum)
			fmt.Println("Your answer is ", j)
		case "*":
			fmt.Println(onenum)
			fmt.Scan(&firstnum)
			fmt.Println(secnum)
			fmt.Scan(&secondnum)
			j := times(firstnum, secondnum)
			fmt.Println("Your answer is ", j)
		case "/":
			fmt.Println(onenum)
			fmt.Scan(&firstnum)
			fmt.Println(secnum)
			fmt.Scan(&secondnum)
			j := divi(firstnum, secondnum)
			fmt.Println("Your answer is ", j)
		default:
			fmt.Println("Not a valid function!")
		}
	}
}
