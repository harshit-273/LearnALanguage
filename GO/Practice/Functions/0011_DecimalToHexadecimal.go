package main

import "fmt"

/*
	- Here we will be creating a function which takes decimal as input and returns hexadecimal.
*/

func decimalToHexadecimal(decNum int) string {
	var deciToHexaMap map[int]rune = map[int]rune{
		0:  '0',
		1:  '1',
		2:  '2',
		3:  '3',
		4:  '4',
		5:  '5',
		6:  '6',
		7:  '7',
		8:  '8',
		9:  '9',
		10: 'A',
		11: 'B',
		12: 'C',
		13: 'D',
		14: 'E',
		15: 'F',
	}
	var hexNum string = ""
	for decNum != 0 {
		hexNum = string(deciToHexaMap[decNum%16]) + hexNum
		decNum /= 16
	}
	return hexNum
}

func main() {
	fmt.Println(decimalToHexadecimal(255))
}
