package main

import "fmt"

/*
	- Here we will be creating a function which takes hexadecimal as input parameter and return decimal of it
*/

func hexadecimalToDecimal(hexNum string) int {
	var hexaToDeciMap map[rune]int = map[rune]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'A': 10,
		'B': 11,
		'C': 12,
		'D': 13,
		'E': 14,
		'F': 15,
	}
	var deciNum int = 0
	var mul16 int = 1

	for i := (len(hexNum) - 1); i >= 0; i-- {
		curr := hexNum[i]
		deciNum = deciNum + (hexaToDeciMap[rune(curr)] * mul16)
		mul16 *= 16
	}
	return deciNum
}

func main() {
	fmt.Println(hexadecimalToDecimal("FF"))
}
