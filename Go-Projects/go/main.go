package main

import "fmt"

func main() {
	var option string
	fmt.Println("What do you want to do?")
	fmt.Println("Options : Make a file (make), Delete a file (del), CD somewhere (cd)")
	fmt.Scan(&option)
}
