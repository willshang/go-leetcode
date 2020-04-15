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

// leetcode234_回文链表
func isPalindrome(head *ListNode) bool {
	m := make([]int, 0)
	for head != nil {
		m = append(m, head.Val)
		head = head.Next
	}
	i, j := 0, len(m)-1
	for i < j {
		if m[i] != m[j] {
			return false
		}
		i++
		j--
	}
	return true
}
