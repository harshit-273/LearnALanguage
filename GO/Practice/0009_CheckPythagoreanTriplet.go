package main

import "fmt"

/*
	- Here we will be checking if the given triplet is pythagorean triplet or not
*/

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func checkPythagoreanTriplet(a int, b int, c int) bool {
	var p int = max(max(a, b), c)
	var q, r int
	if p == a {
		q = b
		r = c
	} else if p == b {
		q = a
		r = c
	} else if p == c {
		q = a
		r = b
	}

	if (p * p) == ((q * q) + (r * r)) {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(checkPythagoreanTriplet(9, 40, 20))
}
