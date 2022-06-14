package main

/*
	Here we will be finding minimum and maximum from the givn array
*/

import "fmt"

func FindMinMax(arr []int) (min int, max int) {
	min = 9999
	max = 0
	for _, value := range arr {
		if min > value {
			min = value
		}
		if max < value {
			max = value
		}
	}
	return
}

func main() {
	fmt.Println(FindMinMax([]int{45, 32, 54, 87, 21, 34}))
}
