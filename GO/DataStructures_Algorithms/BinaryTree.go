package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	data      string
	leftNode  *Node
	rightNode *Node
}

var Root *Node = nil

func CreateNode() (someNode *Node) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the data: ")
	scanner.Scan()
	var input string = scanner.Text()

	someNode = &Node{data: input, leftNode: nil, rightNode: nil}

	return
}

func AddToBinaryTree() {
	newNode := CreateNode()

	if Root == nil {
		Root = newNode
		return
	}
	ptr := Root
	var nodeAdded bool = false
	for !nodeAdded {
		if ptr.data > newNode.data && ptr.leftNode == nil {
			ptr.leftNode = newNode
			nodeAdded = true
		} else if ptr.data <= newNode.data && ptr.rightNode == nil {
			ptr.rightNode = newNode
			nodeAdded = true
		} else if ptr.data > newNode.data && ptr.leftNode != nil {
			ptr = ptr.leftNode
		} else if ptr.data <= newNode.data && ptr.rightNode != nil {
			ptr = ptr.rightNode
		}
	}
}

func DFS_Inorder(node *Node) {
	if Root == nil {
		fmt.Println("Binary Search Tree is empty, enter some data")
		return
	}
	if node == nil {
		return
	}

	DFS_Inorder(node.leftNode)
	fmt.Print(node.data + " ")
	DFS_Inorder(node.rightNode)
}

func DFS_Preorder(node *Node) {
	if Root == nil {
		fmt.Println("Binary Search Tree is empty, enter some data")
		return
	}
	if node == nil {
		return
	}

	fmt.Print(node.data + " ")
	DFS_Preorder(node.leftNode)
	DFS_Preorder(node.rightNode)
}

func DFS_Postorder(node *Node) {
	if Root == nil {
		fmt.Println("Binary Search Tree is empty, enter some data")
		return
	}
	if node == nil {
		return
	}

	DFS_Postorder(node.leftNode)
	DFS_Postorder(node.rightNode)
	fmt.Print(node.data + " ")
}

type Queue struct {
	items [](*Node)
}

func (q *Queue) Enqueue(node *Node) {
	q.items = append(q.items, node)
}

func (q *Queue) Dequeue() *Node {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func BFS_Leverorder() {
	if Root == nil {
		fmt.Println("Binary Search Tree is empty, enter some data")
		return
	}
	myQueue := Queue{}
	myQueue.Enqueue(Root)

	for len(myQueue.items) != 0 {
		current := myQueue.Dequeue()
		fmt.Print(current.data + " ")
		if current.leftNode != nil {
			myQueue.Enqueue(current.leftNode)
		}
		if current.rightNode != nil {
			myQueue.Enqueue(current.rightNode)
		}
	}
}

func Search_BinaryTree(node *Node, searchData string) {
	if node == nil {
		fmt.Print("Data is not present")
		return
	} else {
		if node.data == searchData {
			fmt.Print("Data is present")
			return
		} else if node.data > searchData {
			Search_BinaryTree(node.leftNode, searchData)
		} else {
			Search_BinaryTree(node.rightNode, searchData)
		}
	}
}

func FindMin(node *Node) *Node {
	for node.leftNode != nil {
		node = node.leftNode
	}
	return node
}

func DeleteFromBinaryTree(node *Node, deleteData string) *Node {
	if Root == nil {
		fmt.Println("Nothing to delete, already empty")
	}
	if node == nil {
		return node
	} else if deleteData < node.data {
		node.leftNode = DeleteFromBinaryTree(node.leftNode, deleteData)
	} else if deleteData > node.data {
		node.rightNode = DeleteFromBinaryTree(node.rightNode, deleteData)
	} else {
		if node.leftNode == nil && node.rightNode == nil { // No child
			node = nil
		} else if node.leftNode == nil { // No left or right child
			node = node.rightNode
		} else if node.rightNode == nil {
			node = node.leftNode
		} else { // has both children
			var temp *Node = FindMin(node.rightNode)
			node.data = temp.data
			node.rightNode = DeleteFromBinaryTree(node.rightNode, temp.data)
		}
	}
	return node
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter your choice: ")
	var choice int
	var anyInput string
	var continueWithBST = true
	fmt.Println("Hello! Here you can play with Binary Search Tree as below:")
	for continueWithBST {
		fmt.Println("0 - Exit")
		fmt.Println("1 - Adding data to the Binary Search Tree")
		fmt.Println("2 - Deleting data from the Binary Search Tree")
		fmt.Println("3 - Searching data in the Binary Search Tree")
		fmt.Println("4 - Depth First Search - Inorder")
		fmt.Println("5 - Depth First Search - Preorder")
		fmt.Println("6 - Depth First Search - Postorder")
		fmt.Println("7 - Breadth First Search - Levelorder")
		scanner.Scan()
		choice, _ = strconv.Atoi(scanner.Text())
		switch choice {
		case 0:
			fmt.Println("Exiting ...")
			continueWithBST = false
		case 1:
			AddToBinaryTree()
			fmt.Println()
		case 2:
			fmt.Print("Enter the data to be deleted: ")
			scanner.Scan()
			anyInput = scanner.Text()
			DeleteFromBinaryTree(Root, anyInput)
			fmt.Println()
		case 3:
			fmt.Print("Enter the data to be searched: ")
			scanner.Scan()
			anyInput = scanner.Text()
			Search_BinaryTree(Root, anyInput)
			fmt.Println()
		case 4:
			DFS_Inorder(Root)
			fmt.Println()
		case 5:
			DFS_Preorder(Root)
			fmt.Println()
		case 6:
			DFS_Postorder(Root)
			fmt.Println()
		case 7:
			BFS_Leverorder()
			fmt.Println()
		}
	}
}
