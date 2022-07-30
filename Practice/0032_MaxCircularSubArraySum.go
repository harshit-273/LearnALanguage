package main

import (
	"fmt"
	"math"
)

/*
	- Here we will be finding a maximum sum of circular subarray(think of the array as last element has it's next element as first element)
*/

func MaxCircularSubArraySum(arr []int) (maxSum int) {
	totalSum := 0
	for _, value := range arr {
		totalSum += value
	}
	minSum := math.MaxInt
	curSum := 0
	for _, value := range arr {
		curSum += value
		if curSum > 0 {
			curSum = 0
		}
		minSum = int(math.Min(float64(minSum), float64(curSum)))
	}
	maxSum = totalSum - minSum
	return
}

func main() {
	fmt.Printf("Max Circular SubArray sum is %d\n", MaxCircularSubArraySum([]int{4, -4, 6, -6, 10, -11, 12}))
}
