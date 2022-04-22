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

func main() {
	AddToBinaryTree()
	AddToBinaryTree()
	AddToBinaryTree()
	AddToBinaryTree()
	AddToBinaryTree()
	AddToBinaryTree()
	AddToBinaryTree()

	BFS_Leverorder()
}
