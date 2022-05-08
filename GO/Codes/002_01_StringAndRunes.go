package main

import "fmt"

func main() {
	var someRune rune = '🤪'
	otherRune := "some string"[4]

	fmt.Printf("%c ", someRune)
	fmt.Printf("%d ", otherRune)
}
