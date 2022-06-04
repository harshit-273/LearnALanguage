package main

import "fmt"

/*
	- Here we will be creating a function that takes decimal number as an input and returns binary number.
*/

func decimalToBinary(decNum int) int {
	var binNum int = 0
	var mul10 int = 1
	for decNum != 0 {
		binNum += (decNum % 2) * mul10
		decNum /= 2
		mul10 *= 10
	}
	return binNum
}

func main() {
	fmt.Println(decimalToBinary(5))
}
