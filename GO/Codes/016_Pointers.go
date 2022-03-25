package main

import (
	"fmt"
)

func changeValueWithPointer(str *string) {
	*str = "Value is changed"
}

func changeValueWithoutPointer(str string) {
	str = "This won't work anyway"
}

func main() {
	// assigning and accessing pointer
	x := 7
	y := &x
	fmt.Println(x, *y, y)
	*y = 3
	fmt.Println("After changing the value of *y")
	fmt.Println(x, *y, y)

	// how passing an argument as address(pointer) works
	var someStr string = "Value to be changed"
	var otherStr string = "Other value to be changed"
	changeValueWithPointer(&someStr)
	fmt.Println(someStr)
	changeValueWithoutPointer(otherStr)
	fmt.Println(otherStr)
}
