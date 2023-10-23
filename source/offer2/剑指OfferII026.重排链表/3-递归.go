package main

func main() {
	first := ListNode{Val: 3}
	second := ListNode{Val: 2}
	third := ListNode{Val: 0}
	fourth := ListNode{Val: -4}
	first.Next = &second
	second.Next = &third
	third.Next = &fourth
	fourth.Next = &second
	reorderList(&first)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	helper(head, length)
}

func helper(head *ListNode, length int) *ListNode {
	if length == 1 {
		next := head.Next
		head.Next = nil
		return next
	}
	if length == 2 {
		next := head.Next.Next
		head.Next.Next = nil
		return next
	}
	tail := helper(head.Next, length-2)
	next := tail.Next
	temp := head.Next
	head.Next = tail
	tail.Next = temp
	return next
}
