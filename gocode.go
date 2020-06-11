package main

import (
	"fmt"
	"os"
)

func main() {
	var percentagev string
	var percentaged int
	var percentagef float32
	var leScan string
	var pricey string
	percentagev = "strings"
	percentaged = 69
	percentagef = 3.25
	var getout = "Get out"

	//Section: fmt methods
	fmt.Println("Section: fmt methods")
	fmt.Println("This is the basic print statement : fmt.Println()")
	fmt.Println("It automatically puts a spa", "ce in between segments like the space in between space")
	fmt.Print("This is another print statement but it is fmt.Print()")
	fmt.Print("It does not automatically put a space", "in between segments like spacein and it does not skip a line")
	fmt.Println("\n", "The backslash+n works as a line break when using Print and Println statements.")
	fmt.Println("There are variables like percentagev, percentaged, percentagef that enable you to implement variables in speech. But you have to use Printf. Here is an example of all three below.")
	fmt.Printf("The percentagev is for %v, percentaged is for %d, percentagef is for %.2f, there are a lot of others but these are the basic ones. Try and spot where the percentages could have been used.", percentagev, percentaged, percentagef)
	fmt.Println("For floats if you want it do display the decimal number only up to a certain decimal place you put the % and then .2f with the 2 being the number of decimal places displayed.")
	fmt.Println("The .Scan method is useful for getting the user's answer to a question. you have to make a variable then do fmt.Scan( &(your variable here, no brackets) ).")
	var sprinting = fmt.Sprintf("Sprintf is another method but instead of displaying the text it stores it away in a variable : eg := fmt.Sprintf(). It can then be later used.")
	fmt.Println(sprinting, "The last statement was printed using a Sprintf variable statement.")
	fmt.Println("It is useful when you are working with something like a bank account and you can use .Scan to get the users command which then displays the .Sprintf")
	fmt.Println("The .Scan() method is useful for getting the user's input to a question and the answer can be stored inside a variable for later use.")
	fmt.Println("Type in the magic word to go to the next section(you_can_use_underscores)")
	fmt.Scan(&leScan)

	if leScan != "themagicword" {

		fmt.Println("Wrong! (ツ) ")
		fmt.Scan(&leScan)

	} else {

	}

	if leScan != "themagicword" {
		fmt.Println(getout)
		os.Exit(404)
	}

	//Section : ifs and variables
	fmt.Println("Variables are very important because they are just that ; variables. They can change and can be put anywhere you want(almost).")
	fmt.Println("There are two ways a variable can be declared. one way is to declare the variable but not assign a value to it:", "\n", "var variable1 string = yourstringhereindoublequotes")
	fmt.Println("Or", "\n", "var variable2 = yourvaluehere")
	fmt.Println("Variables come in many types, these are:", "\n", "bool", "\n", "string", "\n", "int", "\n", "uint", "\n", "int8", "\n", "uint8", "\n", "int16", "\n", "uint16", "\n", "int32", "\n", "uint32", "\n", "int64", "\n", "uint64", "\n", "float32", "\n", "float64", "\n", "complex64", "\n", "complex128", "\n", "byte", "\n", "rune", "\n", "error")
	fmt.Println("Now solve this problem:", "\n", "Which one of these three types of variables is most useful for displaying the price of an  item?", "\n", "A) int", "\n", "B) string", "\n", "C) float", "\n", "Type in the letter")
	fmt.Scan(&pricey)
	for pricey != "C" {
		fmt.Println("Wrong!")
		fmt.Scan(&pricey)
	}
	fmt.Println("Correct!")
}

/*

 */
