package main

import (
	"fmt"
	"math"
)

/*
	- Here we will be finding what is the subarray having the highest sum
*/

// Brute Force
func MaxSubArray_BruteForce(arr []int) (maxSum int) {
	lenOfArr := len(arr)
	maxSum = math.MinInt
	for i := 0; i < lenOfArr; i++ {
		for j := i; j < lenOfArr; j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += arr[k]
			}
			if maxSum < sum {
				maxSum = sum
			}
		}
	}
	return
}

// Using Cumulative sum approach
func MaxSubArray_CumulativeSum(arr []int) (maxSum int) {
	lenOfArr := len(arr)
	cumSumArr := make([]int, 0)
	cumSumArr = append(cumSumArr, 0)
	for index, value := range arr {
		cumSumArr = append(cumSumArr, cumSumArr[index]+value)
	}

	lenOfCArr := lenOfArr + 1
	maxSum = math.MinInt
	for i := 1; i < lenOfCArr; i++ {
		for j := 0; j < i; j++ {
			sum := cumSumArr[i] - cumSumArr[j]
			if maxSum < sum {
				maxSum = sum
			}
		}
	}
	return
}

func MaxSubArray_KadaneAlgo(arr []int) (maxSum int) {
	maxSum = math.MinInt
	curSum := 0
	for _, value := range arr {
		curSum += value
		if curSum < 0 {
			curSum = 0
		}
		maxSum = int(math.Max(float64(maxSum), float64(curSum)))
	}
	return
}

func main() {
	fmt.Printf("Maximum sum of subarray is %d\n", MaxSubArray_BruteForce([]int{5, -3, 4, 2}))
	fmt.Printf("Maximum sum of subarray is %d\n", MaxSubArray_CumulativeSum([]int{5, -3, 4, 2}))
	fmt.Printf("Maximum sum of subarray is %d\n", MaxSubArray_KadaneAlgo([]int{-1, 4, -6, 7, -4}))
}
