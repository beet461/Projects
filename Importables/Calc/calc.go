package calc

import "fmt"

func frac(function string, num1, num2 float64) int {
	switch function {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	}
	return 404
}

func calculate(function string, num1 float64, num2 float64) {
	switch function {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	case "%":
		fmt.Println((num1 / 100) * num2)
	case "//":
		frac(function, num1, num2)
	}

}
