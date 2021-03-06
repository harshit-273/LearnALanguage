package main

import "fmt"

func main() {
	////////////////////				Arrays					////////////////////
	/*
		arr := [4]string{"zero", "one", "two", "three"}
		fmt.Printf("%q\n", arr[2])

		var arr2 [2][3]int
		fmt.Printf("Before the assignment of any value to array - %v\n", arr2)

		arr2[1][1] = 4
		fmt.Printf("after the assignment of any value to array - %v\n", arr2)

		fmt.Printf("length of outer array - %d\n", len(arr2))
		fmt.Printf("length of inner array - %d", len(arr2[0]))
	*/

	////////////////////				Slices				////////////////////
	/*
		var sliceArr [5]int = [5]int{0, 1, 2, 3, 4}

		var sliFullCapacity []int = sliceArr[:] // it would give the full array to this slice
		fmt.Println(sliFullCapacity)

		var sli []int = sliceArr[1:4] // this would give sli = {1,2,3} as we started from index 1 and we would end before the index given after ":"
		fmt.Println(sli)
		fmt.Println(cap(sli), len(sli)) // cap function would give how much this slice can store, although the length maybe smaller than or equal to capacity

		otherSli := make([]int, 0) // other way to create a slice without the array
		fmt.Println(otherSli)

		otherSli = append(otherSli, 1) // this would append the element "1" at the end of the slice
		fmt.Println(otherSli)
	*/

	////////////////////				Two Dimensional Slices 				////////////////////

	sliOfSlices := make([][]int, 0) // make a 2D slice of length 0

	fmt.Println(sliOfSlices) // printing the initial slice of slices

	fmt.Println("--------")

	sliOfSlices = append(sliOfSlices, []int{5, 6, 7}) // appending the slice in slice of slices
	sliOfSlices = append(sliOfSlices, []int{1, 2})
	sliOfSlices = append(sliOfSlices, []int{3, 4, 8, 9})

	for outerIndex, sli := range sliOfSlices { // printing the slice of slices with indices and values
		fmt.Printf("outer index: %d =>\n", outerIndex)
		for innerIndex, value := range sli {
			fmt.Printf("inner index: %d -> value: %d, ", innerIndex, value)
		}
		fmt.Println()
	}

	fmt.Println(sliOfSlices) // printing the modified slice of slices
}
