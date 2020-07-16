package main

import "fmt"

func fibonaci(i int) (ret int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}
func main() {
	var i int
	var in int
	fmt.Println("Choose how many numbers of the fibonaci sequence you want to know.")
	fmt.Scan(&in)
	for i = 0; i < in; i++ {
		fmt.Printf("%d ", "\n", " " fibonaci(i))
	}
}
