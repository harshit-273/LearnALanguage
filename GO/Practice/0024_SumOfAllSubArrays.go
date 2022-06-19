package main

import "fmt"

/*
	- Here we will be finding sum of all the subarrays of the array
*/

func SumOfAllSubArray(arr []int) (sol [][]int) {
	lenOfArr := len(arr) - 1
	for i := 0; i <= lenOfArr; i++ {
		sum := 0
		sumArr := make([]int, 0)
		for j := i; j <= lenOfArr; j++ {
			sum += arr[j]
			sumArr = append(sumArr, sum)
		}
		sol = append(sol, sumArr)
	}
	return
}

func main() {
	fmt.Printf("All the sum of the subarrays: %v", SumOfAllSubArray([]int{1, 2, 0, 7, 2}))
}
