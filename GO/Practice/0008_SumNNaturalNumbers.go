package main

import "fmt"

/*
	- Here we will be creating a function which takes input parameter as n and returns sum of first n natural numbers
*/

func sumNNatural(n int) int {
	sum := int(((n * (n + 1)) / 2))
	return sum
}

func main() {
	fmt.Println(sumNNatural(7))
}
