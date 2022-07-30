package main

import "fmt"

/*
	- Here we would be finding a pair which would have the sum equal to k
*/

func PairSum(arr []int, k int) bool {
	iLoop := len(arr) - 1
	jLoop := iLoop + 1
	for i := 0; i < iLoop; i++ {
		for j := 1; j < jLoop; j++ {
			pairSum := arr[i] + arr[j]
			if pairSum == k {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Printf("If any pair was found having sum %d - %t", 31, PairSum([]int{2, 4, 7, 11, 14, 16, 20, 21}, 31))
}
