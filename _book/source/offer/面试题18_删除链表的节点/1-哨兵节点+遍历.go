package main

import "fmt"

func main() {
	first := ListNode{Val: 1}
	second := ListNode{Val: 2}
	third := ListNode{Val: 3}
	first.Next = &second
	second.Next = &third

	deleteNode(&first, 3)
	fmt.Println(first.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 剑指offer_面试题18_删除链表的节点
func deleteNode(head *ListNode, val int) *ListNode {
	headPre := &ListNode{Next: head}
	temp := headPre
	for temp.Next != nil {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
			break
		} else {
			temp = temp.Next
		}
	}
	return headPre.Next
}
