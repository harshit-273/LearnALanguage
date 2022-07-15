package main

import "fmt"

/*
Given an unsorted array A of size N of non-negative integers, find a continuous
subarray which adds to a given number S.
Constraints
1 <= N <= 105
0 <= Ai <= 1010
Example
Input:
N = 5, S = 12
A[] = {1,2,3,7,5}
Output: 2 4
Explanation: The sum of elements from 2nd position to 4th position is 12.
*/

func SubArrayWithGivenSum(arr []int, sum int) (start int, end int) {
	var currSum int = 0
	start = -1
	end = -1
	i := 0
	j := 0
	lenOfArr := len(arr)
	for (j < lenOfArr) && (currSum+arr[j] <= sum) {
		currSum += arr[j]
		j++
	}
	if sum == currSum {
		start = i
		j--
		end = j
		return
	}
	for j < lenOfArr {
		currSum += arr[j]
		for currSum > sum {
			currSum -= arr[i]
			i++
		}
		if sum == currSum {
			start = i
			end = j
			return
		}
		j++
	}
	return
}

func main() {
	st, ed := SubArrayWithGivenSum([]int{1, 2, 3, 7, 5}, 12)
	fmt.Printf("Sub Array with the given sum has staring index:%d and ending index:%d", st, ed)
}
