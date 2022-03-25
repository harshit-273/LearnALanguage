package main

import "fmt"

func main() {
	var someVariable string = fmt.Sprintf("Hello my name is %s and I am %d years old", "Harshit", 22)

	fmt.Printf("%v", someVariable)
}
