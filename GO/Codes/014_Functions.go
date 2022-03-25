package main

import (
	"fmt"
)

func someThing(x int, y int) (int, int) {
	return x + y, x - y
}
func otherThing(x int, y int) (z1 int, z2 int) {
	z1 = x + y
	z2 = x - y
	return
}
func deferExample() {
	defer fmt.Println("Just before closing the function")
	fmt.Println("Other statements")
	fmt.Println("Another statement")
}
func variadicFunction(para ...string) {
	for _, param := range para {
		fmt.Println(param)
	}
}
func init() {
	fmt.Println("init executed")
}

func main() {
	returnedValue1, returnedValue2 := someThing(4, 6)
	fmt.Println(returnedValue1, returnedValue2)

	returnedValue3, returnedValue4 := otherThing(5, 3)
	fmt.Println(returnedValue3, returnedValue4)

	deferExample()

	variadicFunction("one", "two", "three")
}
