package main

import "fmt"

func main() {
	var array1 [12]float64
	var array2 = [2]float64{0.2, 69.9}
	var array3 = [5][3]int{
		{5, 7, 8},
		{6, 8, 7},
		{1, 2, 6},
		{4, 4, 3},
		{7, 8, 7},
	}
	array1[6] = 0.707
	fmt.Scan(&array1[7])
	array1[8] = array1[6] + array1[7]
	fmt.Println(array1[8])
	fmt.Println(array3[0][0] * array3[1][0] * array3[0][1] * array3[1][1] * array3[1][2] * array3[2][1] * array3[2][2])
	fmt.Println(array3[2][2])
	fmt.Println(array2)
}
