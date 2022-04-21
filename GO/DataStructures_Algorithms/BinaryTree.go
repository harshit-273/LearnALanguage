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

func Search_BinaryTree(node *Node, searchData string) (bool, *Node) { // need to modify this by also returning the address of the node having the searchData
	var isPresent bool = false
	var retNode *Node = nil
	if node == nil {
		isPresent = false
	}
	if node.data == searchData {
		isPresent = true
		retNode = node
	} else if node.data > searchData {
		isPresent, retNode = Search_BinaryTree(node.leftNode, searchData)
	} else {
		isPresent, retNode = Search_BinaryTree(node.rightNode, searchData)
	}
	return isPresent, retNode
}

/*	// Still working on this
func RemoveFromBinaryTree(deleteData string) {
	isPresent, removeThisNode := Search_BinaryTree(Root, deleteData)
	if !isPresent {
		fmt.Println("Data is not Present")
		return
	}
	if removeThisNode.leftNode == nil && removeThisNode.rightNode == nil {
		removeThisNode = nil
	} else if removeThisNode.leftNode != nil && removeThisNode.rightNode == nil {
		removeThisNode = removeThisNode.leftNode
	} else if removeThisNode.leftNode == nil && removeThisNode.rightNode != nil {
		removeThisNode = removeThisNode.rightNode
	} else {
		removeThisNode = removeThisNode.rightNode
		var ptr *Node = removeThisNode.rightNode
		for ptr.leftNode != nil {
			ptr = ptr.leftNode
		}
		ptr = removeThisNode.leftNode
	}
}
*/

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
	isPres, val := Search_BinaryTree(Root, input)
	if isPres {
		fmt.Print("Data is present and it's value is " + val.data)
	} else {
		fmt.Print("Data is not present")
	}
}
