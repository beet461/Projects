package main

import (
	"fmt"
	"os"
)

func main() {
	var quest1 string
	var quest2 string
	var quest3 string
	var quest4 string
	var challenge string
	var goo string
	cor := "Correct!"
	fmt.Println("Welcome to the 5 question Grand Quiz, which makes no sense at all!")
	fmt.Println("Rules:", "1. DON'T get it wrong.", "2. Type one character only.", "3. Follow the rules.")
	fmt.Println("Type in -Quiz.- to start the quiz")
	fmt.Scan(&goo)
	if goo != "Quiz." {
		fmt.Println("Nope")
		os.Exit(404)
	}
	fmt.Println("Q1) How does a stone fly?", "Is it: ", "\n", "A) With a hydraulic press", "\n", "B) Chicken Wings", "\n", "C) Lava")
	fmt.Scan(&quest1)
	if quest1 != "Y" {
		fmt.Println("Go Away Moron")
		os.Exit(404)
	} else {
		fmt.Println(cor)
	}
	fmt.Println("Q2) What is an example of a common drink?", "Is it: ", "\n", "A) Coffee", "\n", "C) Air", "\n", "E) Table")
	fmt.Scan(&quest2)
	if quest2 != "E" {
		fmt.Println("Disappointing.")
		os.Exit(404)
	} else {
		fmt.Println(cor)
	}
	fmt.Println("Q3) hwat si two amrledcbs is owt enoqtius hsiT lusp", "ti sI: ", "\n", "E) uofr", "\n", "A)lilmoni ", "\n", "B) ozre")
	fmt.Scan(&quest3)
	if quest3 != "E" {
		fmt.Println("Depressing")
		os.Exit(404)
	} else {
		fmt.Println(cor)
	}
	fmt.Println("Q4) Is this the last question on this quiz (read the first line of this quiz)?", "Is it: ", "\n", "A) oN", "\n", "B) No", "\n", "C) No")
	fmt.Scan(&quest4)
	if quest4 != "T" {
		fmt.Println("Noob")
		os.Exit(404)
	} else {
		fmt.Println(cor)
	}

	fmt.Println("If you are reading this then you have encountered a challenge question!", "\n", "Challenge) What is the answer to to a question?")
	fmt.Println("Is it: ", "\n", "A) 101001010011111001010", "\n", "B) - .... . .-. . / .- .-. . / - .-- --- / - --- ... / .. -. / - .... . / --.- ..- . ... - .. --- -..", "\n", "C) 37456048bvh6s76gwvBVT7IGhg4787yg4")
	fmt.Scan(&challenge)
	if challenge == "answer" {
		fmt.Println("Congratulations, you have a certificate of stupid!")
		os.Exit(398596582049)
	} else if challenge == "........" {
		fmt.Println("Hmm you got most of them right.", "\n", "Now go away nerd.")
		os.Exit(404)
	} else if challenge == "SECRET" {
		fmt.Println("Shh, this is a secret question. Anyway, what is the code word to pass this quiz?")
		var secret string
		fmt.Scan(&secret)
		if secret == "YEET" {
			fmt.Println("YES! You passed the quiz!")
			fmt.Println("Now go away nerd")
			os.Exit(404)
		} else {
			fmt.Println("You're a massive idiot you know")
			os.Exit(404)
		}
	} else {
		fmt.Println("Still wrong.")
		os.Exit(404)
	}
}

/*"Is it: ", "\n", "A) ", "\n", "B) ", "\n", "C) "*/
//copy the line above to make new questions quicker to write
