package main

import (
	"fmt"
	"strconv"
)

func main() {
	intString, _ := strconv.ParseInt("25", 10, 64)
	fmt.Printf("from string to int : %d\n", intString)

	floatString, _ := strconv.ParseFloat("25.52", 64)
	fmt.Printf("from string to float : %f\n", floatString)

	stringInt := strconv.Itoa(12)
	fmt.Printf("from int to string : %q\n", stringInt)

	stringFloat := strconv.FormatFloat(556.54, 'f', 2, 64)
	fmt.Printf("from float to string : %q\n", stringFloat)

	floatInt := float64(65)
	fmt.Printf("from int to float : %f\n", floatInt)

	var someFloat float64 = 56.65
	intFloat := int64(someFloat)
	fmt.Printf("from float to int : %d\n", intFloat)

	var str string = "some😁string"
	runeString := []rune(str)
	fmt.Print("from string to array of runes : ")
	fmt.Printf("%v\n", runeString)

	stringRune := string(runeString)
	fmt.Printf("from array of runes to string : %q\n", stringRune)

}
