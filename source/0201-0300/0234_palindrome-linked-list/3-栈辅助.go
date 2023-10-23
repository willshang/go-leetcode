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
	m := make([]int, 0)
	temp := head
	for temp != nil {
		m = append(m, temp.Val)
		temp = temp.Next
	}
	for head != nil {
		val := m[len(m)-1]
		m = m[:len(m)-1]
		if head.Val != val {
			return false
		}
		head = head.Next
	}
	return true
}
