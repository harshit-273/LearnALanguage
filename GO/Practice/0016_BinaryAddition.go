package main

import "fmt"

/*
	- Here we will be creating a function which takes two binary numbers and returns their addition.
*/

func binaryAddition(num1 int, num2 int) int {
	var carry int = 0
	var ans int = 0
	var mul10 int = 1
	for num1 != 0 || num2 != 0 {
		if num1%10 == 0 && num2%10 == 0 {
			ans = (carry * mul10) + ans
			carry = 0
		} else if (num1%10 != 0 && num2%10 != 1) || (num1%10 != 1 && num2%10 != 0) {
			if carry == 0 {
				ans += mul10
			}
		} else {
			if carry == 0 {
				carry = 1
			} else {
				ans = (carry * mul10) + ans
			}
		}
		num1 /= 10
		num2 /= 10
		mul10 *= 10
	}
	if carry == 1 {
		ans = (carry * mul10) + ans
	}
	return ans
}

func main() {
	fmt.Println(binaryAddition(11, 10))
}
