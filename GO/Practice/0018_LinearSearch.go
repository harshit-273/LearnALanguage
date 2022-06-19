package main

import "fmt"

/*
	- Here we will be finding the key and give it's index in the array using linear search
*/

func LinearSearch(arr []int, key int) (index int) {
	index = -1
	for ind, value := range arr {
		if value == key {
			index = ind
			break
		}
	}
	return
}

func main() {
	var position int = LinearSearch([]int{4, 7, 3, 5, 2, 8, 1, 9}, 6)
	if position != -1 {
		fmt.Printf("position of the key is %d\n", position)
	} else {
		fmt.Println("the key is not present in the array")
	}
}
