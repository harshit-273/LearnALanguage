package main

import "fmt"

/*
	- Here we will be finding the key and give it's index in the array using binary search
	- Here we are assuming that the array passed is sorted. Binary search works on sorted things
*/

func BinarySearch(arr []int, key int) (index int) {
	index = -1
	var start, end, mid int
	start = 0
	end = len(arr) - 1
	for start <= end {
		mid = int((end + start) / 2)
		currMid := arr[mid]
		if currMid == key {
			index = mid
			return
		} else if currMid < key {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return
}

func main() {
	var position int = BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2)
	if position != -1 {
		fmt.Printf("position of the key is %d\n", position)
	} else {
		fmt.Println("the key is not present in the array")
	}
}
