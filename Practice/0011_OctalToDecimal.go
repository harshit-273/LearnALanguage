package main

import "fmt"

/*
	- Here we will be creating a function which takes octal input and gives decimal output
*/

func octalToDecimal(octNum int) int {
	var deciNum int = 0
	var mul8 int = 1
	for octNum != 0 {
		deciNum = deciNum + ((octNum % 10) * mul8)
		mul8 *= 8
		octNum /= 10
	}
	return deciNum
}

func main() {
	fmt.Println(octalToDecimal(101))
}
