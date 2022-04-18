package main

import (
	"bufio"
	"fmt"
	"os"
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
	if node == nil {
		return
	}

	DFS_Inorder(node.leftNode)
	fmt.Print(node.data + " ")
	DFS_Inorder(node.rightNode)
}

func DFS_Preorder(node *Node) {
	if node == nil {
		return
	}

	fmt.Print(node.data + " ")
	DFS_Preorder(node.leftNode)
	DFS_Preorder(node.rightNode)
}

func DFS_Postorder(node *Node) {
	if node == nil {
		return
	}

	DFS_Postorder(node.leftNode)
	DFS_Postorder(node.rightNode)
	fmt.Print(node.data + " ")
}

func Search_BinaryTree(node *Node, searchData string) bool { // need to modify this by also returning the address of the node having the searchData
	var isPresent = false
	if node == nil {
		isPresent = false
	}
	if node.data == searchData {
		isPresent = true
	} else if node.data > searchData {
		Search_BinaryTree(node.leftNode, searchData)
	} else {
		Search_BinaryTree(node.rightNode, searchData)
	}
	return isPresent
}

func RemoveFromBinaryTree(deleteData string) {
	if !Search_BinaryTree(Root, deleteData) {
		return
	}
}

func main() {
	AddToBinaryTree()
	AddToBinaryTree()
	AddToBinaryTree()
	AddToBinaryTree()
	AddToBinaryTree()
	AddToBinaryTree()
	AddToBinaryTree()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the data whose existance has to be checked: ")
	scanner.Scan()
	var input string = scanner.Text()
	if Search_BinaryTree(Root, input) {
		fmt.Print("Data is present")
	} else {
		fmt.Print("Data is not present")
	}

}
