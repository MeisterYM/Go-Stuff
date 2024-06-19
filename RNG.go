package main

//Importing the math package along with fmt.
import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("Computer! What day of the month is it?")

	//Calling the rand method.
	//It should be noted that the math package will look like it disappeared in VS Code.
	//However once you actually use the rand method it will come back.
	//Not sure if that's intentional.
	fmt.Println(rand.IntN(30))
}
