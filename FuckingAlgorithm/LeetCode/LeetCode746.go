package main

import (
	"math"
)

type treeNode746 struct {
	val   int
	index int
	left  *treeNode746
	right *treeNode746
}

func minCostClimbingStairs(cost []int) int {
	ret := math.MaxInt32
	root := BuildTree(0, cost, -1)
	nLength := len(cost)
	InOrder(root, &ret, nLength)
	return ret
}

// InOrder : 遍历树
func InOrder(node *treeNode746, ret *int, nLength int) {
	if node != nil {
		if node.index >= nLength-2 && node.val < *ret && (node.left == nil || node.right == nil) {
			*ret = node.val
		}
		InOrder(node.left, ret, nLength)
		InOrder(node.right, ret, nLength)
	}
}

// BuildTree: 建树
func BuildTree(parentVal int, cost []int, index int) *treeNode746 {
	if index >= len(cost) {
		return nil
	}
	nodeVal := 0
	if index >= 0 {
		nodeVal = cost[index]
	}
	node := &treeNode746{
		val:   nodeVal + parentVal,
		index: index,
		left:  BuildTree(nodeVal, cost, index+1),
		right: BuildTree(nodeVal, cost, index+2),
	}
	return node
}
