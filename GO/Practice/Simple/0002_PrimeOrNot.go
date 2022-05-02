package main

import (
	"fmt"
	"math"
)

/*
	- Here we will be learning how check if a number is prime or not.
	- The number which is only divisible by 1 and itself is a prime number.
*/

func main() {
	var input int = 71
	var isNotPrime bool = false

	if input <= 1 {
		fmt.Println("Please give number greater than 1 as input")
	}
	for i := 2; i <= int(math.Sqrt(float64(input))); i++ {
		if input%i == 0 {
			fmt.Printf("%d is not a prime number", input)
			isNotPrime = true
			break
		}
	}

	if !isNotPrime {
		fmt.Printf("%d is a prime number\n", input)
	}
}
