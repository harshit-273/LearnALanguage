package main

import "fmt"

func main() {
	var num1 int = 5
	var num2 int = 3
	var num3 float32 = 8
	var num4 float32 = 7

	fmt.Printf("%03d + %03d = %03d\n", num1, num2, num1+num2)
	fmt.Printf("%03d - %03d = %03d\n", num1, num2, num1-num2)
	fmt.Printf("%03d * %03d = %03d\n", num1, num2, num1*num2)
	fmt.Printf("%03d / %03d = %03d\n", num1, num2, num1/num2)
	fmt.Printf("%03d %% %03d = %03d\n", num1, num2, num1%num2)
	fmt.Printf("%03.5f / %03.5f = %03.5f\n", num3, num4, num3/num4)

}
