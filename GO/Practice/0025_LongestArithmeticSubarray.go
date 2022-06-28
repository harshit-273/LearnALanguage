package main

import (
	"fmt"
	"math"
)

/***************************************************
Given an array nums of integers, return the length of the longest arithmetic subarray in nums.

Example 1:
Input: nums = [3,6,9,12]
Output: 4
Explanation:
The whole array is an arithmetic array with steps of length = 3.

Example 2:
Input: nums = [9,4,7,10]
Output: 3
Explanation:
The longest arithmetic array is [4,7,10].

Example 3:
Input: nums = [20,1,15,3,10,5,8]
Output: 2
Explanation:
All the subarrays of length 2 are longest as there is no longer subarrays.
***************************************************/

func LongestArithmeticSubarray(arr []int) (sol int) {
	sol = 2
	var lenOfArr int = len(arr) - 1
	prevCommDiff := arr[1] - arr[0]
	currLen := 2
	for i := 2; i <= lenOfArr; i++ {
		if arr[i]-arr[i-1] == prevCommDiff {
			currLen++
		} else {
			prevCommDiff = arr[i] - arr[i-1]
			currLen = 2
		}
		sol = int(math.Max(float64(sol), float64(currLen)))
	}
	return
}

func main() {
	fmt.Printf("Length of the longest subsequence of the array is %v", LongestArithmeticSubarray([]int{20, 1, 15, 3, 10, 5, 8}))
}
