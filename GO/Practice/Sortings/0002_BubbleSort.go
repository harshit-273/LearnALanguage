package main

import "fmt"

/*
	- Here we will be implementing bubble sort
*/

func BubbleSort(unsortedArr []int) []int {
	lenOfArr := len(unsortedArr) - 1
	unsortedArrLen := lenOfArr
	for j := 0; j < lenOfArr; j++ {
		for i := 0; i < unsortedArrLen; i++ {
			if unsortedArr[i] > unsortedArr[i+1] {
				temp := unsortedArr[i]
				unsortedArr[i] = unsortedArr[i+1]
				unsortedArr[i+1] = temp
			}
		}
		unsortedArrLen -= 1
	}
	return unsortedArr
}

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Printf("Before sorting : %v\nAfter sorting : %v", arr, BubbleSort(arr))
}
