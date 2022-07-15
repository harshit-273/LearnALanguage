package main

import "fmt"

/*
	- Here we will be creating a function which takes decimal number as input and returns octal number.
*/

func decimalToOctal(decNum int) int {
	var octNum int = 0
	var mul10 int = 1
	for decNum != 0 {
		octNum += (decNum % 8) * mul10
		decNum /= 8
		mul10 *= 10
	}
	return octNum
}

func main() {
	fmt.Println(decimalToOctal(93))
}
