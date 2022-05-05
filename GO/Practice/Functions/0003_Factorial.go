package main

import "fmt"

/*
	- Here we are going create a function which takes the number whose factorial is to be found and returns the factorial of that number
*/

func factorial(num int) int {
	factorialOfNum := 1
	for i := 2; i <= num; i++ {
		factorialOfNum *= i
	}
	return factorialOfNum
}

func main() {
	fmt.Println(factorial(6))
}
