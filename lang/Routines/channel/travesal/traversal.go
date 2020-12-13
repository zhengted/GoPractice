package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func BuildTree() *Node {
	root := Node{Val: 1}
	root.Left = &Node{2, nil, nil}
	root.Right = &Node{3, nil, nil}
	root.Left.Right = &Node{4, nil, nil}
	root.Right.Left = &Node{5, nil, nil}
	return &root
}

func (n *Node) TraverseFunc(f func(node *Node)) {
	if n != nil {
		n.Left.TraverseFunc(f)
		f(n)
		n.Right.TraverseFunc(f)
	}
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}

func main() {
	root := BuildTree()
	maxNode := 0
	c := root.TraverseWithChannel()
	for node := range c {
		if node.Val > maxNode {
			maxNode = node.Val
		}
	}
	fmt.Printf("Max Node Val:%d\n", maxNode)
}
