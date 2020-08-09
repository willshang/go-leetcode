package main

import "fmt"

func main() {
	first := ListNode{Val: 1}
	firsttwo := ListNode{Val: 2}
	firstthree := ListNode{Val: 2}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	removeDuplicateNodes(&first)
	for {
		fmt.Println(first.Val)
		if first.Next == nil {
			break
		}
		first = *first.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 程序员面试金典02.01_移除重复节点
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	m := make(map[int]bool)
	m[head.Val] = true
	temp := head
	for temp.Next != nil {
		if m[temp.Next.Val] == true {
			temp.Next = temp.Next.Next
		} else {
			m[temp.Next.Val] = true
			temp = temp.Next
		}
	}
	return head
}
