package main

import (
	"fmt"
	"math"
)

/*
Given an array arr[] of size N. The task is to find the first repeating element in an
array of integers, i.e., an element that occurs more than once and whose index of
first occurrence is smallest.
Constraints
1 <= N <= 10^6
0 <= Ai <= 10^6
Example
Input:
7
1 5 3 4 3 5 6
Output:
2
Explanation:
5 is appearing twice and its first appearance is at index 2 which is less than 3
whose first occurring index is 3.
*/

func FirstRepeatingElement(arr []int) (index int) {
	var freq map[int]int = make(map[int]int)
	index = math.MaxInt
	for ind, value := range arr {
		_, ok := freq[value]
		if !ok {
			freq[value] = ind
		} else {
			if index > freq[value] {
				index = freq[value]
			}
		}
	}
	return
}

func main() {
	fmt.Printf("The index of the 1st repeating number is %d", FirstRepeatingElement([]int{1, 5, 3, 4, 3, 5, 6}))
}
