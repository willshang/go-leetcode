package main

import "fmt"

func main() {
	first := ListNode{Val: 3}
	second := ListNode{Val: 2}
	third := ListNode{Val: 0}
	fourth := ListNode{Val: -4}
	first.Next = &second
	second.Next = &third
	third.Next = &fourth

	fmt.Println(insertionSortList(&first))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode147_对链表进行插入排序
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := &ListNode{Next: head}
	cur := head.Next
	head.Next = nil
	for cur != nil {
		next := cur.Next
		prev := res // 从头开始寻找插入点
		for prev.Next != nil && prev.Next.Val <= cur.Val {
			prev = prev.Next
		}
		// 插入操作
		cur.Next = prev.Next
		prev.Next = cur
		// 指向下一个未排序节点
		cur = next
	}
	return res.Next
}
