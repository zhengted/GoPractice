package main

import (
	"GoPractice/lang/FunctionalProgram/tree"
	"fmt"
)

// 闭包实现自定义遍历方法
// 不需要修饰如何访问自由变量

func main() {
	var root tree.TreeNode
	root = tree.TreeNode{Val: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}
	root.Left.Right = &tree.TreeNode{2, nil, nil}
	root.Right.Left = &tree.TreeNode{4, nil, nil}
	root.Traverse()
	count := 0
	root.TraverseFunc(func(node *tree.TreeNode) {
		count++
	})
	fmt.Printf("node count:%d", count)
}
