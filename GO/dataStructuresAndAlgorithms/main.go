package main

import (
	"bufio"
	D "dataStructuresAndAlgorithms/dataStructures"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the data for the node: ")
	var choice int

	var continueWithLinkedList = true
	fmt.Println("Hello! Here you can play with Linked List as below:")
	for continueWithLinkedList {
		fmt.Println("0 - Exit")
		fmt.Println("1 - Create node at the beginning of the Linked List")
		fmt.Println("2 - Create node at the end of the Linked List")
		fmt.Println("3 - Delete node from the brginning of the Linked List")
		fmt.Println("4 - Delete node from the end of the Linked List")
		fmt.Println("99 - Print the data in the whole Linked List")
		scanner.Scan()
		choice, _ = strconv.Atoi(scanner.Text())

		switch choice {
		case 0:
			fmt.Println("Exiting ...")
			continueWithLinkedList = false
		case 1:
			D.CreateBegin()
		case 2:
			D.CreateEnd()
		case 3:
			D.DeleteBegin()
		case 4:
			D.DeleteEnd()
		case 99:
			D.PrintLinkedList()
		}

	}
}
