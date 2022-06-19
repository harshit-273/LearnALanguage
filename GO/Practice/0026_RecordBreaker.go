package main

import "fmt"

/*
Isyana is given the number of visitors at her local theme park on N consecutive days. The number of visitors on the i-th day is Vi. A day is record breaking if it satisfies both of the following conditions:

- The number of visitors on the day is strictly larger than the number of visitors on each of the previous days.
- Either it is the last day, or the number of visitors on the day is strictly larger than the number of visitors on the following day.
Note that the very first day could be a record breaking day!
Please help Isyana find out the number of record breaking days.

Input
The first line of the input gives the number of test cases, T. T test cases follow. Each test case begins with a line containing the integer N. The second line contains N integers. The i-th integer is Vi.

Output
For each test case, output one line containing Case #x: y, where x is the test case number (starting from 1) and y is the number of record breaking days.
*/

func RecordBreaker(arr []int) (sol int) {
	sol = 0
	var max int = -1
	arr = append(arr, -1) // added -1 at the end because we are also comparing next day to our current day
	for index, value := range arr {
		if (value > max) && (value > arr[index+1]) {
			max = value
			sol++
		}
	}
	return
}

func main() {
	fmt.Printf("number of record breaker days are %v", RecordBreaker([]int{1, 2, 0, 7, 2, 0, 2, 2}))
}
