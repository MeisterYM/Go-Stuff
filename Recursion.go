package main

import "fmt"

func main() {
	//Calling the countdown function.
	countdown(-5)
}

//So long as the given number is greater than 0, the function will call itself and decrease said number by 1.
func countdown(number int) int {
	if number < 0 {
		fmt.Println("I can't do a countdown with that number!")
		return 0
	} else if number > 0 {
		fmt.Println(number)
		return countdown(number - 1)
	} else {
		fmt.Println("Liftoff!")
		return 0
	}
}
