package main

import "fmt"

func main() {
	first := ListNode{Val: 2}
	first1 := ListNode{Val: 4}
	//first2 := ListNode{Val: 3}
	first.Next = &first1
	//first1.Next = &first2
	fmt.Println(rotateRight(&first, 1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	temp := head
	count := 0
	arr := make([]*ListNode, 0)
	for temp != nil {
		arr = append(arr, temp)
		temp = temp.Next
		count++
	}
	k = k % count
	if k == 0 {
		return head
	}
	arr[count-1].Next = head
	temp = arr[count-1-k]
	head, temp.Next = temp.Next, nil
	return head
}
