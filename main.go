package main

import "fmt"

func main() {
	drink := "Energy Drink"
	backwardDrink := reversedDrink(drink)
	fmt.Println(backwardDrink)
}

func reversedDrink(input string) (output string) {
	// for loop to reverse
	// End of input working backwards
	// Continue as long as we have characters
	for x := len(input) - 1; x >= 0; x-- {
		output = output + string(input[x])
	}
	return

}
