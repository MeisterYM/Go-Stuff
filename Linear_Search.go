package main

import "fmt"

func main() {

	//Commented code to test variables and methods.
	//fmt.Print("Hello World! This is a test!")

	//The number that the algorithm is going to search for.
	var lookup int = 3

	//The declared array. Using odd numbers to make it interesting. They probably don't need to be sorted either.
	odds := [5]int{1, 3, 5, 7, 9}

	//Search for the value that the variable lookup has. Tell us if you found it and where, otherwise say nothing.
	for i := 0; i < len(odds); i++ {
		if lookup == odds[i] {
			fmt.Println("Found it! It's at index", i)
		}
	}
}
