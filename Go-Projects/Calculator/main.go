package main

import (
	"fmt"

	calc "github.com/beet461/Importable/Calc"
)

func main() {
	var funx string
	var nu1 float64
	var nu2 float64
	fmt.Println("What is your equation (space needed)?")
	fmt.Scan(&nu1, &funx, &nu2)
	calc.Calculate(funx, nu1, nu2)
	main()
}
