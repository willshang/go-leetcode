package main

import "fmt"

func main() {
	first := ListNode{Val: 1}
	firsttwo := ListNode{Val: 2}
	firstthree := ListNode{Val: 2}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	deleteDuplicates(&first)

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

// leetcode82_删除排序链表中的重复元素II
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := &ListNode{Next: head}
	cur := temp
	value := 0
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			value = cur.Next.Val
			for cur.Next != nil && cur.Next.Val == value {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return temp.Next
}
