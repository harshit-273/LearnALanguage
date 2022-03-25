package main

import "fmt"

func main() {
	/*
		Comment the other loop while using one loop
	*/

	/////////////////////////////////               while loop						/////////////////////////////////
	/*
		x := 5
		for x > 0 {
			fmt.Printf("*")
			x--
		}
	*/

	/////////////////////////////////               for loop						/////////////////////////////////
	/*
		for x := 5; x > 0; x-- {
			fmt.Printf("*")
		}
	*/

	/////////////////////////////////              infinite loop					/////////////////////////////////
	/*
		for {
		}
	*/

	/////////////////////////////////              use of "break" and "continue"	/////////////////////////////////
	/*
		printing number between 1-100 that are divisible by 3 or 5, but not both -- continue statement use commented "***1***"
		print the first number divisible by 3 and 5 between 1-100 -- break statement use commented "***2***"
		just remove the comments in front as required
	*/
	/*
		for x := 1; x <= 101; x++ {
			if x%3 == 0 && x%5 == 0 {
				fmt.Printf("%d", x) // ***2***
				break               // ***2***
				// ***1*** continue // ***1***
				// ***1*** } else if (x%3 == 0 && x%5 != 0) || (x%3 != 0 && x%5 == 0) { // ***1***
				// ***1***	fmt.Printf("%4d", x) // ***1***
			}
		}
	*/

	/////////////////////////////////              	use of range					/////////////////////////////////
	arr := [5]int{123, 231, 1232, 12333, 2314}
	for index, element := range arr { // if index is not required then use "_" instead of it
		fmt.Printf("%5d: %5d\n", index, element)
	}
}
