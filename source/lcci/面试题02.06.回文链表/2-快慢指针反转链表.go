package main

import "fmt"

func main() {
	first := new(ListNode)
	first.Val = 1
	second := new(ListNode)
	second.Val = 2
	third := new(ListNode)
	third.Val = 2
	fourth := new(ListNode)
	fourth.Val = 1
	first.Next = second
	second.Next = third
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
