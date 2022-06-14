package main

import "fmt"

func SelectionSort(unsortedArr []int) []int {
	lenOfArr := len(unsortedArr) - 1
	for i := 0; i < lenOfArr; i++ {
		min := i
		for j := (i + 1); j <= lenOfArr; j++ {
			if unsortedArr[min] > unsortedArr[j] {
				min = j
			}
		}
		if min != i {
			temp := unsortedArr[i]
			unsortedArr[i] = unsortedArr[min]
			unsortedArr[min] = temp
		}
	}
	return unsortedArr
}

func main() {
	arr := []int{4, 5, 2, 8, 1, 0, 6, 9, 3, 7}
	fmt.Print("Before sorting : ")
	fmt.Println(arr)
	fmt.Print("After sorting : ")
	sortedArr := SelectionSort(arr)
	fmt.Println(sortedArr)
}
