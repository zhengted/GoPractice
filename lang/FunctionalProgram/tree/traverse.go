package tree

import "fmt"

func (n TreeNode) Traverse() {
	n.TraverseFunc(func(node *TreeNode) {
		fmt.Print(node.Val)
	})
	fmt.Println()
}

func (n *TreeNode) TraverseFunc(f func(node *TreeNode)) {
	if n != nil {
		n.Left.TraverseFunc(f)
		f(n)
		n.Right.TraverseFunc(f)
	}
}
