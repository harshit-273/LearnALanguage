package main

import "fmt"

func main() {
	var someStr string = "h🙂rsHit"
	var arrayOfRune []rune = []rune(someStr)
	fmt.Println(string(someStr[len(someStr) - 3]))
	fmt.Println(arrayOfRune[4])
	fmt.Println(string(arrayOfRune))
}