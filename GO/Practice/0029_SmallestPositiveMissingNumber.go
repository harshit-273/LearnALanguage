package main

import (
	"fmt"
)

/*
Find the smallest positive missing number in the given array.
Example: [0, -10, 1, 3, -20], Ans = 2
*/

func SmallestPositiveMissingNumber(arr []int) (smPsMs int) {
	smPsMs = -1
	var checkMap map[int]bool = make(map[int]bool)
	for _, value := range arr {
		checkMap[value] = true
	}
	for i := 1; i < 99999; i++ {
		if !checkMap[i] {
			smPsMs = i
			return
		}
	}
	return
}

func main() {
	fmt.Printf("Smallest Positive Missing in array is %d", SmallestPositiveMissingNumber([]int{0, 4, -6, 2, 1}))
}
