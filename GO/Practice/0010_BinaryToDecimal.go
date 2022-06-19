package main

import "fmt"

/*
	- Here we will be creating a function which takes binary input and gives decimal output
*/

func binaryToDecimal(binNum int) int {
	var deciNum int = 0
	var mul2 int = 1
	for binNum != 0 {
		deciNum = deciNum + ((binNum % 10) * mul2)
		mul2 *= 2
		binNum /= 10
	}
	return deciNum
}

func main() {
	fmt.Println(binaryToDecimal(1010))
}
