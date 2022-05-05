package main

import "fmt"

/*
	- Here we will be learnig to make a function that takes number of terms and prints that number of fibonacci terms.
*/

func FibonacciSeries(terms int) []int {
	var fibonacci []int = make([]int, 0)
	fibonacci = append(fibonacci, 1)
	fibonacci = append(fibonacci, 1)

	if terms > 2 {
		for i := 2; i < terms; i++ {
			fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
		}
	}
	return fibonacci
}

func main() {
	fmt.Print(FibonacciSeries(5))
}
