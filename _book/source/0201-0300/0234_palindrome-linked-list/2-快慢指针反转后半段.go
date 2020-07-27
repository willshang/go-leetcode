package main

import "fmt"

func main() {
	first := new(ListNode)
	first.Val = 11
	second := new(ListNode)
	second.Val = 20
	second2 := new(ListNode)
	second2.Val = 778
	third := new(ListNode)
	third.Val = 208
	fourth := new(ListNode)
	fourth.Val = 118
	first.Next = second
	second.Next = second2
	second2.Next = third
	third.Next = fourth

	fmt.Println(isPalindrome(first))
	for {
		if first == nil {
			break
		}
		fmt.Print(first.Val, " ")
		first = first.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var pre *ListNode
	cur := slow
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	for pre != nil {
		if head.Val != pre.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	}
	return true
}
