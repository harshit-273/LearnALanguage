package main

import "fmt"

/*
	- Here we need to find max till ith element
*/

func MaxTillI(arr []int) (maxArr []int) {
	var max int = -99999
	for _, value := range arr {
		if max < value {
			max = value
		}
		maxArr = append(maxArr, max)
	}
	return
}

func main() {
	fmt.Printf("Max till i - %v", MaxTillI([]int{-9, 14, 35, 21, 45, 32, 0, 1, 67}))
}
