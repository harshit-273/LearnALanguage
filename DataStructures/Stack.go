package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	data string
	next *Node
}

var Begin *Node = nil

func CreateNode() (someNode *Node) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the data: ")
	scanner.Scan()
	var input string = scanner.Text()

	someNode = &Node{data: input, next: nil}

	return
}

func AddToStack() {
	endNode := CreateNode()

	if Begin == nil {
		Begin = endNode
		return
	}
	ptr := Begin
	for ptr.next != nil {
		ptr = ptr.next
	}
	ptr.next = endNode
}

func RemoveFromStack() {
	if Begin == nil {
		fmt.Println("Stack is already empty")
		return
	}
	if Begin.next == nil {
		Begin = nil
		return
	}
	ptr := Begin
	for ptr.next.next != nil {
		ptr = ptr.next
	}
	ptr.next = nil
}

func PrintStack() {
	if Begin == nil {
		fmt.Println("Nothing to display, Stack is empty.")
		return
	}
	var ptr *Node = Begin

	for ptr != nil {
		fmt.Print("\"" + ptr.data + "\" ")
		ptr = ptr.next
	}
	fmt.Println()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter your choice: ")
	var choice int

	var continueWithStack = true
	fmt.Println("Hello! Here you can play with Queue as below:")
	for continueWithStack {
		fmt.Println("0 - Exit")
		fmt.Println("1 - Add to Stack")
		fmt.Println("2 - Remove from Stack")
		fmt.Println("99 - Print the data in the whole Stack")
		scanner.Scan()
		choice, _ = strconv.Atoi(scanner.Text())

		switch choice {
		case 0:
			fmt.Println("Exiting ...")
			continueWithStack = false
		case 1:
			AddToStack()
		case 2:
			RemoveFromStack()
		case 99:
			PrintStack()
		}
	}
}
