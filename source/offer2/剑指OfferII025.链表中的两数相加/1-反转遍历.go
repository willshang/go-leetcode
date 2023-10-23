package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	first := ListNode{Val: 2}
	first1 := ListNode{Val: 4}
	first2 := ListNode{Val: 3}
	first.Next = &first1
	first1.Next = &first2

	second := ListNode{Val: 5}
	second1 := ListNode{Val: 6}
	second2 := ListNode{Val: 4}
	second.Next = &second1
	second1.Next = &second2

	fmt.Println(addTwoNumbers(&first, &second))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverse(l1)
	l2 = reverse(l2)
	res := &ListNode{}
	cur := res
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10 // 进位
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	return reverse(res.Next)
}

func reverse(head *ListNode) *ListNode {
	var result *ListNode
	var temp *ListNode
	for head != nil {
		temp = head.Next
		head.Next = result
		result = head
		head = temp
	}
	return result
}
