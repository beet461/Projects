package main

import (
	"fmt"
	"os"
)

func make() {
	var path string
	fmt.Println("Please type in your path , e.g. ~/Desktop/...")
	fmt.Scan(&path)
}

func del() {

}

func cd() {

}

func main() {
	var  string
	var option string
	fmt.Println("What do you want to do?")
	fmt.Println("Options : Make a file (make), Delete a file (del), CD somewhere (cd); anything else is to exit.")
	fmt.Scan(&option)
	if option == "make" {
		make()
	} else if option == "del" {
		del()
	} else if option == "cd" {
		cd()
	} else {
		fmt.Println("Do you really want to exit?")
		fmt.Scan(&option)
		if option == "no" {
			main()
		} else {
			os.Exit(410)
		}
	}
}
