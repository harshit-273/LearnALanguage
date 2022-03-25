package main

import "fmt"

func main() {
	///////////////////////////							Normal switch 							///////////////////////////
	/*
		x := 3
		switch x {
		case 3, -3:
			fmt.Printf("three")
		case 5:
			fmt.Printf("five")
		default:
			fmt.Printf("No value")
		}
	*/

	///////////////////////////							If-Else like switch						///////////////////////////
	y := 10
	z := 10
	switch {
	case y < z:
		fmt.Printf("y < z")
	case y > z:
		fmt.Printf("y > z")
	default:
		fmt.Printf("y == z")
	}
}
