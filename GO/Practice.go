package main

import (
	"fmt"
)

func main() {

	sliOfSlices := make([][]int, 4)

	for sliIndex := range sliOfSlices {
		sliOfSlices[sliIndex] = make([]int, 2)
	}

	for _, sli := range sliOfSlices {
		for _, value := range sli {
			fmt.Print(value)
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println("--------")

	sliOfSlices = append(sliOfSlices, []int{5, 6, 7})

	for _, sli := range sliOfSlices {
		for _, value := range sli {
			fmt.Print(value)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
