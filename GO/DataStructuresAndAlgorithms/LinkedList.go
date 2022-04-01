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

var begin *Node = nil

func createNode() (someNode *Node) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the data for the node: ")
	scanner.Scan()
	var input string = scanner.Text()

	someNode = &Node{data: input, next: nil}

	return
}

func createBegin() {
	beginNode := createNode()

	if begin != nil {
		beginNode.next = begin
	}
	begin = beginNode
}

func createEnd() {
	endNode := createNode()
	ptr := begin
	if begin != nil {
		begin = endNode
		return
	}
	for ptr.next != nil {
		ptr = ptr.next
	}
	ptr.next = endNode
}

func printLinkedList() {
	if begin == nil {
		fmt.Println("Nothing to display, Linked List is empty.")
		return
	}
	var ptr *Node = begin

	for ptr != nil {
		fmt.Print("\"" + ptr.data + "\" ")
		ptr = ptr.next
	}
	fmt.Println()
}

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
		fmt.Println("99 - Print the data in the whole Linked List")
		scanner.Scan()
		choice, _ = strconv.Atoi(scanner.Text())

		switch choice {
		case 0:
			fmt.Println("Exiting ...")
			continueWithLinkedList = false
		case 1:
			createBegin()
		case 2:
			createEnd()
		case 99:
			printLinkedList()
		}

	}
}
