package main

import (
	"fmt"
)

func test() {
	fmt.Println("assigning function to a variable")
}

// passing the function as a parameter
func test4(test5 func(int) int) {
	fmt.Println("passing a function as parameter", test5(7))
}

// returning the function from a function
func test7(x string) func() {
	return func() { fmt.Println(x) }
}

func main() {
	// assigning function to a variable
	x := test
	x()

	// declaring a function inside a function
	test2 := func() {
		fmt.Println("declaring a function inside a function")
	}
	test2()

	// calling a function directly and assigning it's value to a variable
	test3 := func(x int) int {
		return x * -1
	}(8)
	fmt.Println("calling a function directly and assigning it to a variable", test3)

	// for explaination of passing function as a parameter
	test6 := func(x int) int {
		return x * 3
	}
	/*
		- first test6 is passed to test4
		- test6 is passed the value "7" in the function test4
		- processing the test4
	*/
	test4(test6)

	// for explaination of returning a function
	/*
		test7 returns a function so it can be assigned to a variable and then it can be called or it can be called directly
	*/
	someFunc := test7("calling a function that returns a function and assigning it to a variable")
	someFunc()

	// directly calling a function that returns a function
	test7("directly calling a function that returns a function")()
}
