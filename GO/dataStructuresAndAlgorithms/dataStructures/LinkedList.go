package dataStructures

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	data string
	next *Node
}

var Begin *Node = nil

func CreateNode() (someNode *Node) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the data for the node: ")
	scanner.Scan()
	var input string = scanner.Text()

	someNode = &Node{data: input, next: nil}

	return
}

func CreateBegin() {
	beginNode := CreateNode()

	if Begin != nil {
		beginNode.next = Begin
	}
	Begin = beginNode
}

func CreateEnd() {
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

func DeleteBegin() {
	beginNode := Begin
	if beginNode == nil {
		fmt.Println("It is Empty")
	}

	Begin = Begin.next
	beginNode = nil
}

func DeleteEnd() {
	if Begin == nil {
		fmt.Println("It is Empty")
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

func PrintLinkedList() {
	if Begin == nil {
		fmt.Println("Nothing to display, It is Empty.")
		return
	}
	var ptr *Node = Begin

	for ptr != nil {
		fmt.Print("\"" + ptr.data + "\" ")
		ptr = ptr.next
	}
	fmt.Println()
}

/*
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
			CreateBegin()
		case 2:
			CreateEnd()
		case 3:
			DeleteBegin()
		case 4:
			DeleteEnd()
		case 99:
			PrintLinkedList()
		}

	}
}*/
