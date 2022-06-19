package main

import "fmt"

/*
	- Here are going to reverse a number. i.e. converting 1234 to 4321
*/

func main() {
	var input int = 1234
	var reversedNumber int = 0

	var newNumber int
	for newNumber = input; newNumber != 0; newNumber = int(newNumber / 10) {
		reversedNumber *= 10
		reversedNumber += newNumber % 10
	}

	fmt.Printf("reverse of number %d is %d", input, reversedNumber)
}
