package main

import "fmt"

func main() {
	first := ListNode{Val:2}
	first1 := ListNode{Val:4}
	first2 := ListNode{Val:3}
	first.Next = &first1
	first1.Next = &first2

	second := ListNode{Val:5}
	second1 := ListNode{Val:6}
	second2 := ListNode{Val:4}
	second.Next = &second1
	second1.Next = &second2

	fmt.Println(addTwoNumbers(&first,&second))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	carry := 0
	for l1 != nil || l2 != nil || carry > 0{
		sum := carry
		if l1 != nil{
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 !=  nil{
			sum += l2.Val
			l2 = l2.Next
		}

		//进位
		carry = sum / 10

		cur.Next = &ListNode{Val:sum%10}
		cur = cur.Next
	}
	return res.Next
}
