package main

import "fmt"

func main() {
	first := ListNode{Val: 1}
	firsttwo := ListNode{Val: 2}
	firstthree := ListNode{Val: 2}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	oddEvenList(&first)

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

func oddEvenList(head *ListNode) *ListNode {
	odd := make([]*ListNode, 0)
	even := make([]*ListNode, 0)
	count := 1
	for head != nil {
		if count%2 == 1 {
			odd = append(odd, head)
		} else {
			even = append(even, head)
		}
		count++
		head = head.Next
	}
	temp := &ListNode{}
	node := temp
	for i := 0; i < len(odd); i++ {
		node.Next = odd[i]
		node = node.Next
	}
	for i := 0; i < len(even); i++ {
		node.Next = even[i]
		node = node.Next
	}
	node.Next = nil
	return temp.Next
}
