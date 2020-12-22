package main

import (
	"fmt"
)

func main() {
	root := BuildTree103()
	fmt.Println(zigzagLevelOrder(root))
}
func BuildTree103() *TreeNode {
	root := TreeNode{Val: 3}
	root.Left = &TreeNode{9, nil, nil}
	root.Right = &TreeNode{20, nil, nil}
	root.Right.Left = &TreeNode{15, nil, nil}
	root.Right.Right = &TreeNode{7, nil, nil}
	return &root
}
func groupAnagramsDemo() {
	param := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagramsEx(param))
}
