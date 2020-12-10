package main

func main() {
	head1 := &ListNode{
		1,
		nil,
	}
	head2 := &ListNode{
		9,
		&ListNode{
			9,
			nil,
		},
	}
	PrintNum(addTwoNumbers(head1, head2))
}
