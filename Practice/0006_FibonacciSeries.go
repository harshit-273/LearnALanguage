package main

import "fmt"

/*
	- Here we will be learnig to make a function that takes number of terms and returns that number of fibonacci terms.
*/

func fibonacciSeries(terms int) []int {
	var fibonacci []int = make([]int, 0)
	fibonacci = append(fibonacci, 0)
	fibonacci = append(fibonacci, 1)

	if terms > 2 {
		for i := 2; i < terms; i++ {
			fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
		}
	}
	return fibonacci
}

func main() {
	fmt.Print(fibonacciSeries(10))
}
