package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	nLen1, nLen2 := CalcLength(*l1), CalcLength(*l2)
	if nLen1 < nLen2 {
		l1, l2 = l2, l1

	}
	head1 := l1

	// 保证l1是较长的那个
	for l1 != nil {
		if l2 != nil {
			l1.Val = l1.Val + l2.Val
		}

		if l1.Val >= 10 {
			l1.Val = l1.Val % 10
			if l1.Next == nil {
				l1.Next = &ListNode{
					Val:  1,
					Next: nil,
				}
				break
			} else {
				l1.Next.Val++
			}
		}
		l1 = l1.Next
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return head1
}

func CalcLength(l ListNode) int {
	res := 1
	for l.Next != nil {
		res++
		l = *(l.Next)
	}
	return res
}

func PrintNum(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
