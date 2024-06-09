package main

import "fmt"

func main() {
	fmt.Println("Uh, this still work?")

	//Calling the string function.
	confirm()

	fmt.Println("Okay but how about doing some math?")

	//Calling the additiong function.
	add(1, 2)
}

//Function for printing a string.
func confirm() {
	fmt.Println("Sure does!")
}

//Function for adding two numbers.
func add(x int, y int) {
	var result = x + y

	fmt.Println(x, "+", y, "=", result, "\nYup I can do that too!")
}
