package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// zigzagLevelOrder: Z字型打印二叉树
// 辅助栈：如果是偶数层，需要从右往左打印
// 辅助队：如果是奇数层，正常的层序遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	nLayer := 1
	var s []*TreeNode
	var totalQ []*TreeNode
	var ret [][]int
	if root == nil {
		return ret
	}
	totalQ = append(totalQ, root)
	totalQ = append(totalQ, nil) // 哨兵元素
	for len(totalQ) > 1 {

		// 常规的层次遍历
		curNode := totalQ[0]
		totalQ = totalQ[1:]
		if curNode == nil {
			// 如果当前层遍历完毕，根据层数决定不同的读取方式
			var tempRet []int // 这里必须新建切片 如果在外部定义 因为切片是数组的映射 会影响原数组
			if nLayer%2 == 1 {
				// 奇数层 顺序
				for len(s) > 0 {
					tempRet = append(tempRet, s[0].Val)
					s = s[1:]
				}
			} else {
				// 偶数层 逆向读取
				for len(s) > 0 {
					tempRet = append(tempRet, s[len(s)-1].Val)
					s = s[:len(s)-1]
				}
			}
			if len(tempRet) > 0 {
				ret = append(ret, tempRet)
			}
			fmt.Println(ret)
			nLayer++
			totalQ = append(totalQ, nil)
			s = s[0:0]
		} else {
			s = append(s, curNode) // 将当前结点插入临时表
			if curNode.Left != nil {
				totalQ = append(totalQ, curNode.Left)
			}
			if curNode.Right != nil {
				totalQ = append(totalQ, curNode.Right)
			}
		}
	}
	var tempRet []int
	if nLayer%2 == 1 {
		// 奇数层 顺序
		for len(s) > 0 {
			tempRet = append(tempRet, s[0].Val)
			s = s[1:]
		}
	} else {
		// 偶数层 逆向读取
		for len(s) > 0 {
			tempRet = append(tempRet, s[len(s)-1].Val)
			s = s[:len(s)-1]
		}
	}
	if len(tempRet) > 0 {
		ret = append(ret, tempRet)
	}
	return ret
}
