package main

import (
	"fmt"
	"math"
)

/***************************************************
Given an array nums of integers, return the length of the longest arithmetic subsequence in nums.

Recall that a subsequence of an array nums is a list nums[i1], nums[i2], ..., nums[ik] with 0 <= i1 < i2 < ... < ik <= nums.length - 1, and that a sequence seq is arithmetic if seq[i+1] - seq[i] are all the same value (for 0 <= i < seq.length - 1).



Example 1:

Input: nums = [3,6,9,12]
Output: 4
Explanation:
The whole array is an arithmetic sequence with steps of length = 3.
Example 2:

Input: nums = [9,4,7,2,10]
Output: 3
Explanation:
The longest arithmetic subsequence is [4,7,10].
Example 3:

Input: nums = [20,1,15,3,10,5,8]
Output: 4
Explanation:
The longest arithmetic subsequence is [20,15,10,5].
***************************************************/

func LongestArithmeticSubsequence(arr []int) (sol int) {
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
	fmt.Printf("Length of the longest subsequence of the array is %v", LongestArithmeticSubsequence([]int{10, 7, 4, 6, 8, 10, 11}))
}
