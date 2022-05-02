package main

import "fmt"

/*
	- Here we will be checking if the given number is a armstrong number. e.g. 153 = 1^3 + 5^3 + 3^3
*/

func main() {
	var input int = 153
	var addedCubes int = 0

	for otherNumber := input; otherNumber != 0; otherNumber = int(otherNumber / 10) {
		cube := otherNumber % 10
		addedCubes += cube * cube * cube
	}

	if addedCubes == input {
		fmt.Printf("%d is an armstrong number", input)
	} else {
		fmt.Printf("%d is not an armstrong nummber", input)
	}
}
