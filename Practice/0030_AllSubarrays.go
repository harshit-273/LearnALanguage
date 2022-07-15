package main

import "fmt"

/*
	- Here will be outputing all the subarrays of the arr.
*/

func AllSubArrays(arr []int) (subarrays [][]int) {
	lenOfArr := len(arr)
	for i := 0; i < lenOfArr; i++ {
		for j := i; j < lenOfArr; j++ {
			subArray := make([]int, 0)
			for k := i; k <= j; k++ {
				subArray = append(subArray, arr[k])
			}
			subarrays = append(subarrays, subArray)
		}
	}
	return
}

func main() {
	fmt.Printf("All the subarrays are as follow:\n%v", AllSubArrays([]int{5, -3, 4, 2}))
}
