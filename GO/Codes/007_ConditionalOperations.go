package main

import "fmt"

func main() {
	var num1 int = 5
	var string1 string = "Harshit"

	fmt.Printf("%d < 5 => %t\n", num1, num1 < 5)
	fmt.Printf("%d > 5 => %t\n", num1, num1 > 5)
	fmt.Printf("%d <= 5 => %t\n", num1, num1 <= 5)
	fmt.Printf("%d >= 5 => %t\n", num1, num1 >= 5)
	fmt.Printf("%d == 5 => %t\n", num1, num1 == 5)
	fmt.Printf("%d != 5 => %t\n\n", num1, num1 != 5)

	fmt.Printf("%q < \"\" => %t\n", string1, string1 < "")
	fmt.Printf("%q > \"\" => %t\n", string1, string1 > "")
	fmt.Printf("%q <= \"\" => %t\n", string1, string1 <= "")
	fmt.Printf("%q >= \"\" => %t\n", string1, string1 >= "")
	fmt.Printf("%q == \"\" => %t\n", string1, string1 == "")
	fmt.Printf("%q != \"\" => %t\n", string1, string1 != "")
}
