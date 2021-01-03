package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func partition(head *ListNode, x int) *ListNode {
	first := &ListNode{}
	p := first
	second := &ListNode{}
	q := second
	for head != nil {
		if head.Val < x {
			if p.Next == nil {
				p.Next = &ListNode{}
			}
			p.Next.Val = head.Val
			p = p.Next
		} else {
			if q.Next == nil {
				q.Next = &ListNode{}
			}
			q.Next.Val = head.Val
			q = q.Next
		}
		head = head.Next
	}
	p.Next = second.Next
	return first.Next
}
