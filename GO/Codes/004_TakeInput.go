package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter something : ")
	scanner.Scan()
	var input string = scanner.Text()
	fmt.Printf("You have typed %q\n", input)

	// another way of taking inputs
	fmt.Printf("Please enter some text:")
	fmt.Scanf("%s", &input)
	fmt.Printf("you have entered %q", input)
}
