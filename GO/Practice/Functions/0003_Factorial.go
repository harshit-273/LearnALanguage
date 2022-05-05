package main

import "fmt"

/*
	- Here we are going create a function which takes the number whose factorial is to be found and returns the factorial of that number

	- Here other function that is implemented is nCr(combination) input is n and r and returns it's nCr value

	- Here one other function is implemented is pascal's triangle:
	01
	01 01
	01 02 01
	01 03 03 01
	01 04 06 04 01
	01 05 10 10 05 01
*/

func factorial(num int) int {
	factorialOfNum := 1
	if num == 0 || num == 1 {
		factorialOfNum = 1
	}
	for i := 2; i <= num; i++ {
		factorialOfNum *= i
	}
	return factorialOfNum
}

func nCr(n int, r int) int {
	var valueOf_nCr float64 = float64(factorial(n)) / (float64(factorial(r)) * float64(factorial(n-r)))
	return int(valueOf_nCr)
}

func pascalTriangle(n int) {
	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%02d ", nCr(i, j))
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println(factorial(6))
	fmt.Println(nCr(6, 4))
	pascalTriangle(5)
}
