package main

import (
	"fmt"
	"os"
)

func make(path) {
	fmt.Println("Please type in your path , e.g. ~/Desktop/...")
	fmt.Scan(&path)
}

func del(path) {

}

func cd(path) {

}

func main() {
	var path string
	var option string
	fmt.Println("What do you want to do?")
	fmt.Println("Options : Make a file (make), Delete a file (del), CD somewhere (cd); anything else is to exit.")
	fmt.Scan(&option)
	if option == "make" {
		make(path)
	} else if option == "del" {
		del(path)
	} else if option == "cd" {
		cd(path)
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
