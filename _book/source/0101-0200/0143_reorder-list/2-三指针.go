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

// leetcode143_重排链表
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	cur, prev, next := head, head, head
	for cur != nil {
		next = cur.Next
		// prev 指向n-1 prev.next 指向n
		for prev = next; prev != nil && prev.Next != nil && prev.Next.Next != nil; {
			prev = prev.Next
		}
		if prev != nil && prev.Next != nil {
			cur.Next = prev.Next
			prev.Next.Next = next
			prev.Next = nil
		}
		cur = next
	}
}
