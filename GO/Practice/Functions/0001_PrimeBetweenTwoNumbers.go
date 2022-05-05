package main

import (
	"fmt"
	"math"
)

/*
	- Here we will be learning how to create a function that determines if the number is prime or not.
	- We will be using this functioin to print the prime numbers between two numbers(both numbers included)
*/

func isPrime(num int) bool {
	sqrtOfNum := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrtOfNum; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var beginNum int = 2
	var endNum int = 45

	for i := beginNum; i <= endNum; i++ {
		if isPrime(i) {
			fmt.Printf("%03d ", i)
		}
	}
}
