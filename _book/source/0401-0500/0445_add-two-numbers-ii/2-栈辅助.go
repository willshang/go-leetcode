package main

import "fmt"

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

// leetcode445_两数相加II
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := make([]int, 0)
	stack2 := make([]int, 0)
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	var res *ListNode
	carry := 0
	for len(stack1) > 0 || len(stack2) > 0 || carry > 0 {
		if len(stack1) > 0 {
			carry = carry + stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			carry = carry + stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		temp := &ListNode{
			Val:  carry % 10,
			Next: res,
		}
		carry = carry / 10
		res = temp
	}
	return res
}
