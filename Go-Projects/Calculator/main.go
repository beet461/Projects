package main

import (
	"fmt"

	calc "github.com/beet461/Importable/Calc"
)

func main() {
	var funx string
	var nu1 float64
	var nu2 float64
	fmt.Println("Choose your function, +, -, *, /, %, //(fraction) or exit.")
	fmt.Scan(&funx)
	fmt.Println("Choose your numbers (space needed).")
	fmt.Scan(&nu1, &nu2)
	calc.Calculate(funx, nu1, nu2)
}
