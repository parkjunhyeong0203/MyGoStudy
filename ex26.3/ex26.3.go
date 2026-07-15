package main

import "fmt"

type Node[T any] struct {
	val  T
	next *Node[T]
}

func NewNode[T any](v T) *Node[T] {
	return &Node[T]{val: v}
}

func (n *Node[T]) Push(v T) *Node[T] {
	node := NewNode(v)
	n.next = node
	return node
}

func main() {
	node1 := NewNode(1)
	node1.Push(2).Push(3).Push(4)

	for node1 != nil {
		fmt.Print(node1.val, "-")
		node1 = node1.next
	}

	fmt.Println()
	node2 := NewNode("hi")
	node2.Push("Hello").Push("How are you").Push("fine, and you?")

	for node2 != nil {
		fmt.Print(node2.val, "-")
		node2 = node2.next
	}
}
