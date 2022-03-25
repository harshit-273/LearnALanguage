package main

import "fmt"

func main() {
	fmt.Printf("true || true => %t\n", true || true)
	fmt.Printf("true || false => %t\n", true || false)
	fmt.Printf("false || true => %t\n", false || true)
	fmt.Printf("false || false => %t\n\n", false || false)

	fmt.Printf("true && true => %t\n", true && true)
	fmt.Printf("true && false => %t\n", true && false)
	fmt.Printf("false && true => %t\n", false && true)
	fmt.Printf("false && false => %t\n\n", false && false)

	fmt.Printf("!true => %t\n", !true)
	fmt.Printf("!false => %t\n", !false)
}
