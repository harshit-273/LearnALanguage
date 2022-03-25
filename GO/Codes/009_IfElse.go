package main

import "fmt"

func main() {
	var condition bool = false
	var other_condition bool = false

	if condition == true && other_condition == false {
		fmt.Printf("condition is %t, but other_condition is %t so \"if\" block is executed", condition, other_condition)
	} else if condition == false && other_condition == true {
		fmt.Printf("condition is %t, but other_condition is %t so \"else if\" block is executed", condition, other_condition)
	} else {
		fmt.Printf("condition is %t and other_condition is also %t so \"else\" block is executed", condition, other_condition)
	}
}
