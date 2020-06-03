package main

import "fmt"

func main() {
	first := ListNode{Val: 1}
	second := ListNode{Val: 3}
	third := ListNode{Val: 2}
	first.Next = &second
	second.Next = &third
	fmt.Println(reversePrint(&first))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 剑指offer_面试题06_从尾到头打印链表
func reversePrint(head *ListNode) []int {
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	i := 0
	for i < len(res)/2 {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
		i++
	}
	return res
}
