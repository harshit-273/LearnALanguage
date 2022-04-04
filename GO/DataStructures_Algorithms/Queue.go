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

func AddToQueue() {
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

func RemoveFromQueue() {
	beginNode := Begin
	if beginNode == nil {
		fmt.Println("Queue is already empty")
		return
	}
	Begin = Begin.next
	beginNode = nil
}

func PrintQueue() {
	if Begin == nil {
		fmt.Println("Nothing to display, Queue is empty.")
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

	var continueWithQueue = true
	fmt.Println("Hello! Here you can play with Queue as below:")
	for continueWithQueue {
		fmt.Println("0 - Exit")
		fmt.Println("1 - Add to Queue")
		fmt.Println("2 - Remove from Queue")
		fmt.Println("99 - Print the data in the whole Queue")
		scanner.Scan()
		choice, _ = strconv.Atoi(scanner.Text())

		switch choice {
		case 0:
			fmt.Println("Exiting ...")
			continueWithQueue = false
		case 1:
			AddToQueue()
		case 2:
			RemoveFromQueue()
		case 99:
			PrintQueue()
		}

	}
}
