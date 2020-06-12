package main

import (
	"fmt"
	"os"
)

func quiz() {
	var quest1 string
	var quest2 string
	var quest3 string
	var goo string
	cor := "Correct!"
	fmt.Println("Welcome to the Grand Quiz, which makes no sense at all!")
	fmt.Println("Rules:", "1. DON'T get it wrong.", "2. Type one character only.", "3. Follow the rules.")
	fmt.Println("Type in -Quiz.- to start the quiz")
	fmt.Scan(&goo)
	if goo != "Quiz." {
		fmt.Println("Nope")
		os.Exit(404)
	}
	fmt.Println("Q1) How does a stone fly?", "Is it: ", "\n", "A) With a hydraulic press", "\n", "B) Chicken Wings", "\n", "C) Lava")
	fmt.Scan(&quest1)
	if quest1 != "D" {
		fmt.Println("Go Away Moron")
		os.Exit(404)
	} else {
		fmt.Println(cor)
	}
	fmt.Println("Q2) What is an example of a common drink?", "Is it: ", "\n", "A) Coffee", "\n", "B) Air", "\n", "C) Table")
	fmt.Scan(&quest2)
	if quest2 != "C" {
		fmt.Println("Disappointing.")
		os.Exit(404)
	} else {
		fmt.Println(cor)
	}
	fmt.Println("Q3) hwat si two amrledcbs is owt enoqtius hsiT lusp", "ti sI: ", "\n", "C) uofr", "\n", "A)lilmoni ", "\n", "B) ozre")
	fmt.Scan(&quest3)
}

/*"Is it: ", "\n", "A) ", "\n", "B) ", "\n", "C) "*/
//copy the line above to make new questions quicker to write
