package main

import "fmt"

/*
	- Here we will be implementing insertion sort
*/

func InsertionSort(unsortedArr []int) []int {
	lenOfArr := len(unsortedArr) - 1
	for i := 1; i <= lenOfArr; i++ {
		current := unsortedArr[i]
		j := i - 1
		for j >= 0 && unsortedArr[j] > current {
			unsortedArr[j+1] = unsortedArr[j]
			j--
		}
		unsortedArr[j+1] = current
	}
	return unsortedArr
}

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Printf("Before sorting : %v\nAfter sorting : %v", arr, InsertionSort(arr))
}
